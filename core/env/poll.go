/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"strconv"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/constants"
)

// MaxPollIntervalForSidecar returns the maximum interval for polling by the
// sidecar process. The value is read from the environment variable
// `VSECM_SIDECAR_MAX_POLL_INTERVAL` or returns 300000 milliseconds as default.
func MaxPollIntervalForSidecar() time.Duration {
	p := constants.GetEnv(constants.VSecMSidecarMaxPollInterval)
	d, _ := strconv.Atoi(string(constants.VSecMSidecarMaxPollIntervalDefault))
	if p == "" {
		p = string(constants.VSecMSidecarMaxPollIntervalDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}

// ExponentialBackoffMultiplierForSidecar returns the multiplier for exponential
// backoff by the sidecar process.
// The value is read from the environment variable
// `VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` or returns 2 as default.
func ExponentialBackoffMultiplierForSidecar() int64 {
	p := constants.GetEnv(constants.VSecMSidecarExponentialBackoffMultiplier)
	d, _ := strconv.Atoi(string(constants.VSecMSidecarExponentialBackoffMultiplierDefault))
	if p == "" {
		p = string(constants.VSecMSidecarExponentialBackoffMultiplierDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return i
	}

	return i
}

// SuccessThresholdForSidecar returns the number of consecutive successful
// polls before reducing the interval. The value is read from the environment
// variable `VSECM_SIDECAR_SUCCESS_THRESHOLD` or returns 3 as default.
func SuccessThresholdForSidecar() int64 {
	p := constants.GetEnv(constants.VSecMSidecarSuccessThreshold)
	d, _ := strconv.Atoi(string(constants.VSecMSidecarSuccessThresholdDefault))
	if p == "" {
		p = string(constants.VSecMSidecarSuccessThresholdDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return i
	}

	return i
}

// ErrorThresholdForSidecar returns the number of consecutive failed polls
// before increasing the interval. The value is read from the environment
// variable `VSECM_SIDECAR_ERROR_THRESHOLD` or returns 2 as default.
func ErrorThresholdForSidecar() int64 {
	p := constants.GetEnv(constants.VSecMSidecarErrorThreshold)
	d, _ := strconv.Atoi(string(constants.VSecMSidecarErrorThresholdDefault))
	if p == "" {
		p = string(constants.VSecMSidecarErrorThresholdDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return i
	}

	return i
}

// PollIntervalForSidecar returns the polling interval for sentry in time.Duration
// The interval is determined by the VSECM_SIDECAR_POLL_INTERVAL environment
// variable, with a default value of 20000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func PollIntervalForSidecar() time.Duration {
	p := constants.GetEnv(constants.VSecMSidecarPollInterval)
	d, _ := strconv.Atoi(string(constants.VSecMSidecarPollIntervalDefault))
	if p == "" {
		p = string(constants.VSecMSidecarPollIntervalDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}
