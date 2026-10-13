module github.com/omniboost/go-24sevenoffice

go 1.18

require (
	github.com/cydev/zero v0.0.0-20160322155811-4a4535dd56e7
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pkg/errors v0.9.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require github.com/hashicorp/errwrap v1.0.0 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
