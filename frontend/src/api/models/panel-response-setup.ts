/* tslint:disable */
/* eslint-disable */
/**
 * Ocserv User management Example Api
 * This is a sample Ocserv User management Api server.
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


// May contain unused imports in some cases
// @ts-ignore
import type { ModelsUser } from './models-user';

/**
 * 
 * @export
 * @interface PanelResponseSetup
 */
export interface PanelResponseSetup {
    /**
     * 
     * @type {string}
     * @memberof PanelResponseSetup
     */
    'token': string;
    /**
     * 
     * @type {ModelsUser}
     * @memberof PanelResponseSetup
     */
    'user': ModelsUser;
}

