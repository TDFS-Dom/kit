// Code generated by protoc-gen-go-errors-i18n. DO NOT EDIT.

package v1

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	i18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func ErrorInvalidCredentialsLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidCredentials",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_CREDENTIALS.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_CREDENTIALS.String(), "")
	}

}

func ErrorInvalidOperationLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidOperation",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_OPERATION.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_OPERATION.String(), "")
	}

}

func ErrorUserLockedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "UserLocked",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_USER_LOCKED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_USER_LOCKED.String(), "")
	}

}

func ErrorEmailNotConfirmedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "EmailNotConfirmed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_EMAIL_NOT_CONFIRMED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_EMAIL_NOT_CONFIRMED.String(), "")
	}

}

func ErrorPhoneNotConfirmedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "PhoneNotConfirmed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_PHONE_NOT_CONFIRMED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_PHONE_NOT_CONFIRMED.String(), "")
	}

}

func ErrorEmailRecoverFailedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "EmailRecoverFailed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_EMAIL_RECOVER_FAILED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_EMAIL_RECOVER_FAILED.String(), "")
	}

}

func ErrorEmailConfirmFailedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "EmailConfirmFailed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_EMAIL_CONFIRM_FAILED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_EMAIL_CONFIRM_FAILED.String(), "")
	}

}

func ErrorPhoneRecoverFailedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "PhoneRecoverFailed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_PHONE_RECOVER_FAILED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_PHONE_RECOVER_FAILED.String(), "")
	}

}

func ErrorPhoneConfirmFailedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "PhoneConfirmFailed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_PHONE_CONFIRM_FAILED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_PHONE_CONFIRM_FAILED.String(), "")
	}

}

func ErrorTwoStepFailedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "TwoStepFailed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_TWO_STEP_FAILED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_TWO_STEP_FAILED.String(), "")
	}

}

func ErrorConfirmPasswordMismatchLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "ConfirmPasswordMismatch",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), msg)
	} else {
		return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), "")
	}

}

func ErrorRememberTokenNotFoundLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "RememberTokenNotFound",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_REMEMBER_TOKEN_NOT_FOUND.String(), msg)
	} else {
		return errors.New(403, ErrorReason_REMEMBER_TOKEN_NOT_FOUND.String(), "")
	}

}

func ErrorRememberTokenUsedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "RememberTokenUsed",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_REMEMBER_TOKEN_USED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_REMEMBER_TOKEN_USED.String(), "")
	}

}

func ErrorUserDeletedLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "UserDeleted",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_USER_DELETED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_USER_DELETED.String(), "")
	}

}

func ErrorRefreshTokenInvalidLocalized(localizer *i18n.Localizer, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "RefreshTokenInvalid",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(401, ErrorReason_REFRESH_TOKEN_INVALID.String(), msg)
	} else {
		return errors.New(401, ErrorReason_REFRESH_TOKEN_INVALID.String(), "")
	}

}
