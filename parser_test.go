package main

import (
	"reflect"
	"testing"
)

func TestLineParser(t *testing.T) {

	type test struct {
		line    string
		extract lintResult
	}

	tests := []test{
		{
			"/tmp/test/utils/rest-token-provider.go:54::warning: duplicate of rest-token-provider.go:78-101 (dupl)",
			lintResult{
				"/tmp/test/utils/rest-token-provider.go",
				54,
				0,
				"warning",
				"duplicate of rest-token-provider.go:78-101",
				"dupl",
				"",
			},
		},
		{
			"/tmp/test/utils/config_test.go:58::warning: cyclomatic complexity 11 of function TestLoadRegistryConfiguration() is high (> 10) (gocyclo)",
			lintResult{
				"/tmp/test/utils/config_test.go",
				58,
				0,
				"warning",
				"cyclomatic complexity 11 of function TestLoadRegistryConfiguration() is high (> 10)",
				"gocyclo",
				"",
			},
		},
		{
			"/tmp/test/web/router.go:42:15:error: NewOauthMw not declared by package middleware (gotype)",
			lintResult{
				"/tmp/test/web/router.go",
				42,
				15,
				"error",
				"NewOauthMw not declared by package middleware",
				"gotype",
				"",
			},
		},
		{
			"/tmp/test/utils/token_cache.go:143:3:warning: should use for range instead of for { select {} } (gosimple)",
			lintResult{
				"/tmp/test/utils/token_cache.go",
				143,
				3,
				"warning",
				"should use for range instead of for { select {} }",
				"gosimple",
				"",
			},
		},
		{
			"/private/tmp/test/utils/rest-token-provider.go:90:8:warning: ineffectual assignment to err (ineffassign)",
			lintResult{
				"/private/tmp/test/utils/rest-token-provider.go",
				90,
				8,
				"warning",
				"ineffectual assignment to err",
				"ineffassign",
				"",
			},
		},
		{
			"/tmp/test/middleware/oauth-authentication.go:116:22:error: cannot use ag.tokenCache.GetAuthenticationSummary(token) (value of type *bitbucket.org/eliocity/api-gateway/vendor/bitbucket.org/eliocity/go-common/model/auth.AuthenticationSummary) as *bitbucket.org/eliocity/go-common/model/auth.AuthenticationSummary value in assignment (gotype)",
			lintResult{
				"/tmp/test/middleware/oauth-authentication.go",
				116,
				22,
				"error",
				"cannot use ag.tokenCache.GetAuthenticationSummary(token) (value of type *bitbucket.org/eliocity/api-gateway/vendor/bitbucket.org/eliocity/go-common/model/auth.AuthenticationSummary) as *bitbucket.org/eliocity/go-common/model/auth.AuthenticationSummary value in assignment",
				"gotype",
				"",
			},
		},
	}

	for _, c := range tests {

		l, err := parseLine(c.line)

		if err != nil {
			t.Fatalf("Regexp does not match to :%s", c.line)
		}

		if !reflect.DeepEqual(l, &c.extract) {
			t.Errorf("Mismatch '%s' is != '%s'", l, c.extract)
		}
	}

}
