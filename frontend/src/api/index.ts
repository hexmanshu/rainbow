// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import Api from "./model";

const url = import.meta.env.VITE_API_URL as string || "http://localhost:8080"

const api = new Api({
  url: url,
  verbose: true,
});

export default api;