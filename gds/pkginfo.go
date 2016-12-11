
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Data structures, version 0.1.0.
package gds

import "gitlab.com/vbit/semver"

var version semver.Version

func init() {
    version = semver.New(0,1,0)
}

func Version() semver.Version {
    return version
}

