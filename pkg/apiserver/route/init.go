package route

import (
    "github.com/elrondwong/multistack-example/pkg/apiserver/route/tencents3"
)

func init() {
    InstallRouteFuncs = append(InstallRouteFuncs, tencents3.AddRoute)
}