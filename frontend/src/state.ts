// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'
import { User } from "@snowind/state";

const breakpoints = useBreakpoints(breakpointsTailwind);

const isMobile = breakpoints.smaller('sm');
const isTablet = breakpoints.between('sm', 'lg');
const isDesktop = breakpoints.greater('lg');

const user = new User();

export { user, isMobile, isTablet, isDesktop }