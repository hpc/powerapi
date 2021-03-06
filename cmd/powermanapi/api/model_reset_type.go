/*
 * Power API v1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lowell@lanl.gov
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type ResetType string

// List of ResetType
const (
	RESETTYPE_ON ResetType = "On"
	RESETTYPE_FORCE_OFF ResetType = "ForceOff"
	RESETTYPE_GRACEFUL_SHUTDOWN ResetType = "GracefulShutdown"
	RESETTYPE_GRACEFUL_RESTART ResetType = "GracefulRestart"
	RESETTYPE_FORCE_RESTART ResetType = "ForceRestart"
	RESETTYPE_NMI ResetType = "Nmi"
	RESETTYPE_FORCE_ON ResetType = "ForceOn"
	RESETTYPE_PUSH_POWER_BUTTON ResetType = "PushPowerButton"
	RESETTYPE_POWER_CYCLE ResetType = "PowerCycle"
)
