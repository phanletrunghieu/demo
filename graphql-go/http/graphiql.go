package http

import "net/http"

type graphiQLHandler struct{}

// ServeHTTP .
func (*graphiQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("only GET requests are supported"))
		return
	}

	w.Write(graphiql)
}

// NewGraphiQLHandler .
func NewGraphiQLHandler() http.Handler {
	return &graphiQLHandler{}
}

var graphiql = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.css"/>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
		<style>
			body {
				height: 100%;
				margin: 0;
				width: 100%;
				overflow: hidden;
			}
			#graphiql {
				height: 100vh;
			}
			.jwt-token {
				background: linear-gradient(#f7f7f7, #e2e2e2);
				border-bottom: 1px solid #d0d0d0;
				font-family: system, -apple-system, 'San Francisco', '.SFNSDisplay-Regular', 'Segoe UI', Segoe, 'Segoe WP', 'Helvetica Neue', helvetica, 'Lucida Grande', arial, sans-serif;
				padding: 7px 14px 6px;
				font-size: 14px;
			}
		</style>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div class="jwt-token">JWT Token <input id="jwt-token" placeholder="JWT Token goes here"></div>
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			var search = window.location.search;
			var parameters = {};
			document.getElementById('jwt-token').value = localStorage.getItem('graphiql:jwtToken');
			
			function fetchGQL(params) {
				const jwtToken = document.getElementById('jwt-token').value;
				let headers = {
					'Accept': 'application/json',
					'Content-Type': 'application/json'
				};
				if (jwtToken) {
					localStorage.setItem('graphiql:jwtToken', jwtToken);
					headers = {
						'Accept': 'application/json',
						'Content-Type': 'application/json',
						'Authorization': jwtToken ? 'Bearer ' + jwtToken : null
					};
				}

				return fetch("/graphql", {
					method: "post",
					headers,
					body: JSON.stringify(params),
					credentials: "include",
				}).then(function (resp) {
					return resp.text();
				}).then(function (body) {
					try {
						return JSON.parse(body);
					} catch (error) {
						return body;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: fetchGQL}),
				document.getElementById("graphiql")
			)
		</script>
	</body>
</html>
`)
