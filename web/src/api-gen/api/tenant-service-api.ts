/* tslint:disable */
/* eslint-disable */
/**
 * Saas Service
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import {
  DUMMY_BASE_URL,
  assertParamExists,
  setApiKeyToObject,
  setBasicAuthToObject,
  setBearerAuthToObject,
  setOAuthToObject,
  setSearchParams,
  serializeDataIfNeeded,
  toPathString,
  createRequestFunction,
} from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { RpcStatus } from '../models';
// @ts-ignore
import { Tenantv1Tenant } from '../models';
// @ts-ignore
import { V1CreateTenantRequest } from '../models';
// @ts-ignore
import { V1DeleteTenantReply } from '../models';
// @ts-ignore
import { V1ListTenantReply } from '../models';
// @ts-ignore
import { V1ListTenantRequest } from '../models';
// @ts-ignore
import { V1UpdateTenantRequest } from '../models';
/**
 * TenantServiceApi - axios parameter creator
 * @export
 */
export const TenantServiceApiAxiosParamCreator = function (configuration?: Configuration) {
  return {
    /**
     *
     * @param {V1CreateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceCreateTenant: async (
      body: V1CreateTenantRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('tenantServiceCreateTenant', 'body', body);
      const localVarPath = `/v1/saas/tenant`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceDeleteTenant: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('tenantServiceDeleteTenant', 'id', id);
      const localVarPath = `/v1/saas/tenant/{id}`.replace(
        `{${'id'}}`,
        encodeURIComponent(String(id)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} idOrName
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceGetTenant: async (
      idOrName: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'idOrName' is not null or undefined
      assertParamExists('tenantServiceGetTenant', 'idOrName', idOrName);
      const localVarPath = `/v1/saas/tenant/{idOrName}`.replace(
        `{${'idOrName'}}`,
        encodeURIComponent(String(idOrName)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {string} [filterNameLike]
     * @param {Array<string>} [filterRegionIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceListTenant: async (
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterNameLike?: string,
      filterRegionIn?: Array<string>,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      const localVarPath = `/v1/saas/tenants`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      if (pageOffset !== undefined) {
        localVarQueryParameter['pageOffset'] = pageOffset;
      }

      if (pageSize !== undefined) {
        localVarQueryParameter['pageSize'] = pageSize;
      }

      if (search !== undefined) {
        localVarQueryParameter['search'] = search;
      }

      if (sort) {
        localVarQueryParameter['sort'] = sort;
      }

      if (fields !== undefined) {
        localVarQueryParameter['fields'] = fields;
      }

      if (filterIdIn) {
        localVarQueryParameter['filter.idIn'] = filterIdIn;
      }

      if (filterNameIn) {
        localVarQueryParameter['filter.nameIn'] = filterNameIn;
      }

      if (filterNameLike !== undefined) {
        localVarQueryParameter['filter.nameLike'] = filterNameLike;
      }

      if (filterRegionIn) {
        localVarQueryParameter['filter.regionIn'] = filterRegionIn;
      }

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {V1ListTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceListTenant2: async (
      body: V1ListTenantRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('tenantServiceListTenant2', 'body', body);
      const localVarPath = `/v1/saas/tenant/list`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceUpdateTenant: async (
      tenantId: string,
      body: V1UpdateTenantRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'tenantId' is not null or undefined
      assertParamExists('tenantServiceUpdateTenant', 'tenantId', tenantId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('tenantServiceUpdateTenant', 'body', body);
      const localVarPath = `/v1/saas/tenant/{tenant.id}`.replace(
        `{${'tenant.id'}}`,
        encodeURIComponent(String(tenantId)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceUpdateTenant2: async (
      tenantId: string,
      body: V1UpdateTenantRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'tenantId' is not null or undefined
      assertParamExists('tenantServiceUpdateTenant2', 'tenantId', tenantId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('tenantServiceUpdateTenant2', 'body', body);
      const localVarPath = `/v1/saas/tenant/{tenant.id}`.replace(
        `{${'tenant.id'}}`,
        encodeURIComponent(String(tenantId)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'PATCH', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
  };
};

/**
 * TenantServiceApi - functional programming interface
 * @export
 */
export const TenantServiceApiFp = function (configuration?: Configuration) {
  const localVarAxiosParamCreator = TenantServiceApiAxiosParamCreator(configuration);
  return {
    /**
     *
     * @param {V1CreateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceCreateTenant(
      body: V1CreateTenantRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Tenantv1Tenant>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceCreateTenant(
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceDeleteTenant(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1DeleteTenantReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceDeleteTenant(
        id,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} idOrName
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceGetTenant(
      idOrName: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Tenantv1Tenant>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceGetTenant(
        idOrName,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {string} [filterNameLike]
     * @param {Array<string>} [filterRegionIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceListTenant(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterNameLike?: string,
      filterRegionIn?: Array<string>,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListTenantReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceListTenant(
        pageOffset,
        pageSize,
        search,
        sort,
        fields,
        filterIdIn,
        filterNameIn,
        filterNameLike,
        filterRegionIn,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {V1ListTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceListTenant2(
      body: V1ListTenantRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListTenantReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceListTenant2(
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceUpdateTenant(
      tenantId: string,
      body: V1UpdateTenantRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Tenantv1Tenant>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceUpdateTenant(
        tenantId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async tenantServiceUpdateTenant2(
      tenantId: string,
      body: V1UpdateTenantRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Tenantv1Tenant>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.tenantServiceUpdateTenant2(
        tenantId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
  };
};

/**
 * TenantServiceApi - factory interface
 * @export
 */
export const TenantServiceApiFactory = function (
  configuration?: Configuration,
  basePath?: string,
  axios?: AxiosInstance,
) {
  const localVarFp = TenantServiceApiFp(configuration);
  return {
    /**
     *
     * @param {V1CreateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceCreateTenant(
      body: V1CreateTenantRequest,
      options?: any,
    ): AxiosPromise<Tenantv1Tenant> {
      return localVarFp
        .tenantServiceCreateTenant(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceDeleteTenant(id: string, options?: any): AxiosPromise<V1DeleteTenantReply> {
      return localVarFp
        .tenantServiceDeleteTenant(id, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} idOrName
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceGetTenant(idOrName: string, options?: any): AxiosPromise<Tenantv1Tenant> {
      return localVarFp
        .tenantServiceGetTenant(idOrName, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {string} [filterNameLike]
     * @param {Array<string>} [filterRegionIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceListTenant(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterNameLike?: string,
      filterRegionIn?: Array<string>,
      options?: any,
    ): AxiosPromise<V1ListTenantReply> {
      return localVarFp
        .tenantServiceListTenant(
          pageOffset,
          pageSize,
          search,
          sort,
          fields,
          filterIdIn,
          filterNameIn,
          filterNameLike,
          filterRegionIn,
          options,
        )
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {V1ListTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceListTenant2(
      body: V1ListTenantRequest,
      options?: any,
    ): AxiosPromise<V1ListTenantReply> {
      return localVarFp
        .tenantServiceListTenant2(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceUpdateTenant(
      tenantId: string,
      body: V1UpdateTenantRequest,
      options?: any,
    ): AxiosPromise<Tenantv1Tenant> {
      return localVarFp
        .tenantServiceUpdateTenant(tenantId, body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} tenantId
     * @param {V1UpdateTenantRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    tenantServiceUpdateTenant2(
      tenantId: string,
      body: V1UpdateTenantRequest,
      options?: any,
    ): AxiosPromise<Tenantv1Tenant> {
      return localVarFp
        .tenantServiceUpdateTenant2(tenantId, body, options)
        .then((request) => request(axios, basePath));
    },
  };
};

/**
 * Request parameters for tenantServiceCreateTenant operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceCreateTenantRequest
 */
export interface TenantServiceApiTenantServiceCreateTenantRequest {
  /**
   *
   * @type {V1CreateTenantRequest}
   * @memberof TenantServiceApiTenantServiceCreateTenant
   */
  readonly body: V1CreateTenantRequest;
}

/**
 * Request parameters for tenantServiceDeleteTenant operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceDeleteTenantRequest
 */
export interface TenantServiceApiTenantServiceDeleteTenantRequest {
  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceDeleteTenant
   */
  readonly id: string;
}

/**
 * Request parameters for tenantServiceGetTenant operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceGetTenantRequest
 */
export interface TenantServiceApiTenantServiceGetTenantRequest {
  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceGetTenant
   */
  readonly idOrName: string;
}

/**
 * Request parameters for tenantServiceListTenant operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceListTenantRequest
 */
export interface TenantServiceApiTenantServiceListTenantRequest {
  /**
   *
   * @type {number}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly pageOffset?: number;

  /**
   *
   * @type {number}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly pageSize?: number;

  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly search?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly sort?: Array<string>;

  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly fields?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly filterIdIn?: Array<string>;

  /**
   *
   * @type {Array<string>}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly filterNameIn?: Array<string>;

  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly filterNameLike?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof TenantServiceApiTenantServiceListTenant
   */
  readonly filterRegionIn?: Array<string>;
}

/**
 * Request parameters for tenantServiceListTenant2 operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceListTenant2Request
 */
export interface TenantServiceApiTenantServiceListTenant2Request {
  /**
   *
   * @type {V1ListTenantRequest}
   * @memberof TenantServiceApiTenantServiceListTenant2
   */
  readonly body: V1ListTenantRequest;
}

/**
 * Request parameters for tenantServiceUpdateTenant operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceUpdateTenantRequest
 */
export interface TenantServiceApiTenantServiceUpdateTenantRequest {
  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceUpdateTenant
   */
  readonly tenantId: string;

  /**
   *
   * @type {V1UpdateTenantRequest}
   * @memberof TenantServiceApiTenantServiceUpdateTenant
   */
  readonly body: V1UpdateTenantRequest;
}

/**
 * Request parameters for tenantServiceUpdateTenant2 operation in TenantServiceApi.
 * @export
 * @interface TenantServiceApiTenantServiceUpdateTenant2Request
 */
export interface TenantServiceApiTenantServiceUpdateTenant2Request {
  /**
   *
   * @type {string}
   * @memberof TenantServiceApiTenantServiceUpdateTenant2
   */
  readonly tenantId: string;

  /**
   *
   * @type {V1UpdateTenantRequest}
   * @memberof TenantServiceApiTenantServiceUpdateTenant2
   */
  readonly body: V1UpdateTenantRequest;
}

/**
 * TenantServiceApi - object-oriented interface
 * @export
 * @class TenantServiceApi
 * @extends {BaseAPI}
 */
export class TenantServiceApi extends BaseAPI {
  /**
   *
   * @param {TenantServiceApiTenantServiceCreateTenantRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceCreateTenant(
    requestParameters: TenantServiceApiTenantServiceCreateTenantRequest,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceCreateTenant(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceDeleteTenantRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceDeleteTenant(
    requestParameters: TenantServiceApiTenantServiceDeleteTenantRequest,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceDeleteTenant(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceGetTenantRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceGetTenant(
    requestParameters: TenantServiceApiTenantServiceGetTenantRequest,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceGetTenant(requestParameters.idOrName, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceListTenantRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceListTenant(
    requestParameters: TenantServiceApiTenantServiceListTenantRequest = {},
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceListTenant(
        requestParameters.pageOffset,
        requestParameters.pageSize,
        requestParameters.search,
        requestParameters.sort,
        requestParameters.fields,
        requestParameters.filterIdIn,
        requestParameters.filterNameIn,
        requestParameters.filterNameLike,
        requestParameters.filterRegionIn,
        options,
      )
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceListTenant2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceListTenant2(
    requestParameters: TenantServiceApiTenantServiceListTenant2Request,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceListTenant2(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceUpdateTenantRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceUpdateTenant(
    requestParameters: TenantServiceApiTenantServiceUpdateTenantRequest,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceUpdateTenant(requestParameters.tenantId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {TenantServiceApiTenantServiceUpdateTenant2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof TenantServiceApi
   */
  public tenantServiceUpdateTenant2(
    requestParameters: TenantServiceApiTenantServiceUpdateTenant2Request,
    options?: AxiosRequestConfig,
  ) {
    return TenantServiceApiFp(this.configuration)
      .tenantServiceUpdateTenant2(requestParameters.tenantId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }
}
