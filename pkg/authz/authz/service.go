package authz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/authn"
	"github.com/goxiaoy/go-saas/common"
	"strings"
)

type Service interface {
	CheckForSubjects(ctx context.Context, resource Resource, action Action, subject ...Subject) (*Result, error)
	Check(ctx context.Context, resource Resource, action Action) (*Result, error)
	CheckInTenant(ctx context.Context, resource Resource, action Action, tenantId string) (*Result, error)
}

// SubjectContributor receive one Subject and retrieve as list of subjects
type SubjectContributor interface {
	Process(ctx context.Context, subject Subject) ([]Subject, error)
}

type SubjectResolver interface {
	ResolveFromContext(ctx context.Context) ([]Subject, error)
	ResolveProcessed(ctx context.Context, subjects ...Subject) ([]Subject, error)
}

type SubjectResolverImpl struct {
	opt *Option
}

var _ SubjectResolver = (*SubjectResolverImpl)(nil)

func NewSubjectResolver(opt *Option) *SubjectResolverImpl {
	return &SubjectResolverImpl{opt: opt}
}

func (s *SubjectResolverImpl) ResolveFromContext(ctx context.Context) ([]Subject, error) {
	var subjects []Subject
	var userId string
	userInfo, _ := authn.FromUserContext(ctx)
	userId = userInfo.GetId()
	//append empty user
	subjects = append(subjects, NewUserSubject(userId))
	if clientId, ok := authn.FromClientContext(ctx); ok && len(clientId) > 0 {
		//do not append empty client
		subjects = append(subjects, NewClientSubject(clientId))
	}
	return subjects, nil
}

func (s *SubjectResolverImpl) ResolveProcessed(ctx context.Context, subjects ...Subject) ([]Subject, error) {
	//use contributor
	var subjectList []Subject

	addIfNotPresent := func(subject Subject) bool {
		for _, s := range subjectList {
			if s.GetIdentity() == subject.GetIdentity() {
				return false
			}
		}
		subjectList = append(subjectList, subject)
		return true
	}
	for _, s := range subjects {
		addIfNotPresent(s)
	}
	i := 0
	for {
		if i == len(subjectList) {
			break
		}
		for _, contributor := range s.opt.SubjectContributorList {
			if subjects, err := contributor.Process(ctx, subjectList[i]); err != nil {
				return nil, err
			} else {
				for _, s2 := range subjects {
					addIfNotPresent(s2)
				}
			}
		}
		i++
	}
	return subjectList, nil
}

type Option struct {
	SubjectContributorList []SubjectContributor
}

func NewAuthorizationOption(subjectContributorList ...SubjectContributor) *Option {
	return &Option{SubjectContributorList: subjectContributorList}
}

type DefaultAuthorizationService struct {
	checker PermissionChecker
	sr      SubjectResolver
	log     *log.Helper
}

var _ Service = (*DefaultAuthorizationService)(nil)

func NewDefaultAuthorizationService(checker PermissionChecker, sr SubjectResolver, logger log.Logger) *DefaultAuthorizationService {
	return &DefaultAuthorizationService{checker: checker, sr: sr, log: log.NewHelper(log.With(logger, "module", "authz.service"))}
}

func (a *DefaultAuthorizationService) CheckForSubjects(ctx context.Context, resource Resource, action Action, subject ...Subject) (*Result, error) {
	ti, _ := common.FromCurrentTenant(ctx)
	return a.check(ctx, resource, action, ti.GetId(), subject...)
}

func (a *DefaultAuthorizationService) Check(ctx context.Context, resource Resource, action Action) (*Result, error) {
	ti, _ := common.FromCurrentTenant(ctx)
	return a.CheckInTenant(ctx, resource, action, ti.GetId())
}

func (a *DefaultAuthorizationService) CheckInTenant(ctx context.Context, resource Resource, action Action, tenantId string) (*Result, error) {
	subjects, err := a.sr.ResolveFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return a.check(ctx, resource, action, tenantId, subjects...)
}

func (a *DefaultAuthorizationService) check(ctx context.Context, resource Resource, action Action, tenant string, subject ...Subject) (*Result, error) {
	if always, ok := FromAlwaysAuthorizationContext(ctx); ok {
		var subjectStr []string
		for _, s := range subject {
			subjectStr = append(subjectStr, s.GetIdentity())
		}
		if always {
			a.log.Debugf("check permission for Subject %s Action %s to Resource %s granted", strings.Join(subjectStr, ","), action.GetIdentity(), fmt.Sprintf("%s/%s", resource.GetNamespace(), resource.GetIdentity()))
			return NewAllowAuthorizationResult(), nil
		} else {
			a.log.Debugf("check permission for Subject %s Action %s to Resource %s forbidden", strings.Join(subjectStr, ","), action.GetIdentity(), fmt.Sprintf("%s/%s", resource.GetNamespace(), resource.GetIdentity()))
			r := NewDisallowAuthorizationResult(nil)
			return r, FormatError(ctx, r)
		}
	}

	subjectList, err := a.sr.ResolveProcessed(ctx, subject...)
	if err != nil {
		return nil, err
	}

	var logStr []string
	for _, s := range subjectList {
		logStr = append(logStr, s.GetIdentity())
	}

	logItems := []interface{}{
		strings.Join(logStr, ","), action.GetIdentity(), fmt.Sprintf("%s/%s", resource.GetNamespace(), resource.GetIdentity()), tenant,
	}
	a.log.Debugf("check permission for Subject: %s Action: %s to Resource: %s in Tenant:%s", logItems...)

	grantType, err := a.checker.IsGrantTenant(ctx, resource, action, tenant, subjectList...)
	if err != nil {
		return nil, err
	}
	if grantType == EffectForbidden {
		a.log.Debugf("check permission for Subject: %s Action: %s to Resource: %s in Tenant: %s forbidden", logItems...)
		r := NewDisallowAuthorizationResult(nil)
		return r, FormatError(ctx, r, subjectList...)
	}
	if grantType == EffectGrant {
		a.log.Debugf("check permission for Subject: %s Action: %s to Resource: %s in Tenant: %s granted", logItems...)
		return NewAllowAuthorizationResult(), nil
	}
	a.log.Debugf("check permission for Subject: %s Action: %s to Resource: %s in Tenant: %s forbidden", logItems...)
	r := NewDisallowAuthorizationResult(nil)
	return r, FormatError(ctx, r, subjectList...)
}

var ProviderSet = wire.NewSet(NewDefaultAuthorizationService,
	wire.Bind(new(Service), new(*DefaultAuthorizationService)),
	NewSubjectResolver, wire.Bind(new(SubjectResolver), new(*SubjectResolverImpl)))
