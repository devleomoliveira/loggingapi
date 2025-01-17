package controller

import "net/http"

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(
			`<html>
			<head>
				<title>Logging Server</title>
			</head>
			<body>
				<h1>Hello Logging</h1>
				<ul>
					<li><a href="/swagger/index.html">swagger</a></li>
					<li><a href="/metrics">metrics</a></li>
					<li><a href="/healthz">healthz</a></li>
					<li><a href="/">api list</a></li>
				</ul>
				<hr>
				<center>Logging/0.0</center>
			</body>
		</html>`))
	}
}
