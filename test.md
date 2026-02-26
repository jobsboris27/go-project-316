time="2026-02-26T13:30:13Z" level=warning msg="Found orphan containers ([source-app-run-aec0edfd3a58]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up."
Container source-server-1 Running
Container source-app-1 Creating
Container source-app-1 Created
Attaching to app-1, server-1
app-1 | go test -v ./tests
app-1 | === RUN TestAnalyzeJSONReportSingleDepth
app-1 | analyze_test.go:141: expected 1 page in report, got 3
app-1 | --- FAIL: TestAnalyzeJSONReportSingleDepth (0.00s)
app-1 | === RUN TestCLIIntegrationAgainstFixtures
app-1 | === RUN TestCLIIntegrationAgainstFixtures/simple.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/example.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/blog.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -71,3 +71,3 @@
app-1 | (string) (len=9) "has_title": (bool) true,
app-1 | - (string) (len=5) "title": (string) (len=12) "Crawler Blog"
app-1 | + (string) (len=5) "title": (string) (len=10) "Робот"
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/blog.test
app-1 | Messages: crawler output for http://blog.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/spa.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/missing.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":interface {}(nil), "broken_links":interface {}(nil), "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -5,4 +5,6 @@
app-1 | (map[string]interface {}) (len=9) {
app-1 | - (string) (len=6) "assets": (interface {}) <nil>,
app-1 | - (string) (len=12) "broken_links": (interface {}) <nil>,
app-1 | + (string) (len=6) "assets": ([]interface {}) {
app-1 | + },
app-1 | + (string) (len=12) "broken_links": ([]interface {}) {
app-1 | + },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | Test: TestCLIIntegrationAgainstFixtures/missing.test
app-1 | Messages: crawler output for http://missing.test must match fixture
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures (2.14s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/simple.test (1.67s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/example.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/blog.test (0.08s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/spa.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/missing.test (0.24s)
app-1 | FAIL
app-1 | FAIL project 2.145s
app-1 | FAIL
app-1 | make: \*\*\* [Makefile:35: test] Error 1

app-1 exited with code 2
Aborting on container exit...
Container source-app-1 Stopping
Container source-app-1 Stopped
Container source-server-1 Stopping
Container source-server-1 Stopped

Error: The tests have failed. Examine what they have to say. Inhale deeply. Exhale. Fix the code.
Error: The process '/usr/bin/docker' failed with exit code 2
at ExecState.\_setResult (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:206396)
at ExecState.CheckComplete (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:205956)
at ChildProcess.<anonymous> (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:204797)
1s
0s

ДОПОЛНИЛ!!
/usr/bin/docker compose -f docker-compose.yml up --abort-on-container-exit
time="2026-02-26T13:40:00Z" level=warning msg="Found orphan containers ([source-app-run-4d063fdc4ed8]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up."
Container source-server-1 Running
Container source-app-1 Creating
Container source-app-1 Created
Attaching to app-1, server-1
app-1 | go test -v ./tests
app-1 | === RUN TestAnalyzeJSONReportSingleDepth
app-1 | analyze_test.go:141: expected 1 page in report, got 3
app-1 | --- FAIL: TestAnalyzeJSONReportSingleDepth (0.00s)
app-1 | === RUN TestCLIIntegrationAgainstFixtures
app-1 | === RUN TestCLIIntegrationAgainstFixtures/simple.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/example.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/blog.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -71,3 +71,3 @@
app-1 | (string) (len=9) "has_title": (bool) true,
app-1 | - (string) (len=5) "title": (string) (len=12) "Crawler Blog"
app-1 | + (string) (len=5) "title": (string) (len=10) "Робот"
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/blog.test
app-1 | Messages: crawler output for http://blog.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/spa.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/missing.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":interface {}(nil), "broken_links":interface {}(nil), "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -5,4 +5,6 @@
app-1 | (map[string]interface {}) (len=9) {
app-1 | - (string) (len=6) "assets": (interface {}) <nil>,
app-1 | - (string) (len=12) "broken_links": (interface {}) <nil>,
app-1 | + (string) (len=6) "assets": ([]interface {}) {
app-1 | + },
app-1 | + (string) (len=12) "broken_links": ([]interface {}) {
app-1 | + },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | Test: TestCLIIntegrationAgainstFixtures/missing.test
app-1 | Messages: crawler output for http://missing.test must match fixture
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures (2.07s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/simple.test (1.62s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/example.test (0.07s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/blog.test (0.08s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/spa.test (0.07s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/missing.test (0.22s)
app-1 | FAIL
app-1 | FAIL project 2.068s
app-1 | FAIL
app-1 | make: \*\*\* [Makefile:35: test] Error 1

app-1 exited with code 2
Aborting on container exit...
Container source-app-1 Stopping
Container source-app-1 Stopped
Container source-server-1 Stopping
Container source-server-1 Stopped

Error: The tests have failed. Examine what they have to say. Inhale deeply. Exhale. Fix the code.
Error: The process '/usr/bin/docker' failed with exit code 2
at ExecState.\_setResult (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:206396)
at ExecState.CheckComplete (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:205956)
at ChildProcess.<anonymous> (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:204797)
1s
0s

ДОПОЛНИЛ ЕЩНЕ

#15 resolving provenance for metadata file
#15 DONE 0.0s
cd code && go mod download
/usr/bin/docker compose -f docker-compose.yml up --abort-on-container-exit
time="2026-02-26T13:40:00Z" level=warning msg="Found orphan containers ([source-app-run-4d063fdc4ed8]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up."
Container source-server-1 Running
Container source-app-1 Creating
Container source-app-1 Created
Attaching to app-1, server-1
app-1 | go test -v ./tests
app-1 | === RUN TestAnalyzeJSONReportSingleDepth
app-1 | analyze_test.go:141: expected 1 page in report, got 3
app-1 | --- FAIL: TestAnalyzeJSONReportSingleDepth (0.00s)
app-1 | === RUN TestCLIIntegrationAgainstFixtures
app-1 | === RUN TestCLIIntegrationAgainstFixtures/simple.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/example.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/blog.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -71,3 +71,3 @@
app-1 | (string) (len=9) "has_title": (bool) true,
app-1 | - (string) (len=5) "title": (string) (len=12) "Crawler Blog"
app-1 | + (string) (len=5) "title": (string) (len=10) "Робот"
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/blog.test
app-1 | Messages: crawler output for http://blog.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/spa.test
app-1 | === RUN TestCLIIntegrationAgainstFixtures/missing.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":interface {}(nil), "broken_links":interface {}(nil), "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -5,4 +5,6 @@
app-1 | (map[string]interface {}) (len=9) {
app-1 | - (string) (len=6) "assets": (interface {}) <nil>,
app-1 | - (string) (len=12) "broken_links": (interface {}) <nil>,
app-1 | + (string) (len=6) "assets": ([]interface {}) {
app-1 | + },
app-1 | + (string) (len=12) "broken_links": ([]interface {}) {
app-1 | + },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | Test: TestCLIIntegrationAgainstFixtures/missing.test
app-1 | Messages: crawler output for http://missing.test must match fixture
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures (2.07s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/simple.test (1.62s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/example.test (0.07s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/blog.test (0.08s)
app-1 | --- PASS: TestCLIIntegrationAgainstFixtures/spa.test (0.07s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/missing.test (0.22s)
app-1 | FAIL
app-1 | FAIL project 2.068s
app-1 | FAIL
app-1 | make: \*\*\* [Makefile:35: test] Error 1

app-1 exited with code 2
Aborting on container exit...
Container source-app-1 Stopping
Container source-app-1 Stopped
Container source-server-1 Stopping
Container source-server-1 Stopped

Error: The tests have failed. Examine what they have to say. Inhale deeply. Exhale. Fix the code.
Error: The process '/usr/bin/docker' failed with exit code 2
at ExecState.\_setResult (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:206396)
at ExecState.CheckComplete (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:205956)
at ChildProcess.<anonymous> (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:204797)
1s
0s
#15 DONE 0.0s
cd code && go mod download
/usr/bin/docker compose -f docker-compose.yml up --abort-on-container-exit
time="2026-02-26T14:31:02Z" level=warning msg="Found orphan containers ([source-app-run-5b13eaa70bbf]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up."
Container source-server-1 Running
Container source-app-1 Creating
Container source-app-1 Created
Attaching to app-1, server-1
app-1 | go test -v ./tests
app-1 | === RUN TestAnalyzeJSONReportSingleDepth
app-1 | analyze_test.go:141: expected 1 page in report, got 3
app-1 | --- FAIL: TestAnalyzeJSONReportSingleDepth (0.00s)
app-1 | === RUN TestCLIIntegrationAgainstFixtures
app-1 | === RUN TestCLIIntegrationAgainstFixtures/simple.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Simple Test Site"}, "status":"ok", "url":"http://simple.test"}}, "root_url":"http://simple.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Simple Test Site"}, "status":"ok", "url":"http://simple.test"}}, "root_url":"http://simple.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,7 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | - (string) (len=6) "assets": ([]interface {}) {
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | + (map[string]interface {}) (len=6) {
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | Test: TestCLIIntegrationAgainstFixtures/simple.test
app-1 | Messages: crawler output for http://simple.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/example.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":234, "status_code":200, "type":"image", "url":"http://example.test/assets/img/hero.svg"}, map[string]interface {}{"size_bytes":194, "status_code":200, "type":"script", "url":"http://example.test/assets/site.js"}, map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test"}, "status":"ok", "url":"http://example.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"About Example Test"}, "status":"ok", "url":"http://example.test/about.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test News"}, "status":"ok", "url":"http://example.test/news.html"}}, "root_url":"http://example.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":234, "status_code":200, "type":"image", "url":"http://example.test/assets/img/hero.svg"}, map[string]interface {}{"size_bytes":194, "status_code":200, "type":"script", "url":"http://example.test/assets/site.js"}, map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test"}, "status":"ok", "url":"http://example.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"About Example Test"}, "status":"ok", "url":"http://example.test/about.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test News"}, "status":"ok", "url":"http://example.test/news.html"}}, "root_url":"http://example.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=3) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=3) {
app-1 | @@ -25,4 +25,2 @@
app-1 | },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -40,3 +38,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -48,4 +46,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -64,3 +60,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -72,4 +68,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/example.test
app-1 | Messages: crawler output for http://example.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/blog.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=5) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=2) {
app-1 | @@ -19,4 +19,2 @@
app-1 | },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -34,3 +32,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -42,4 +40,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -58,7 +54,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | - (string) (len=6) "assets": ([]interface {}) {
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | + (map[string]interface {}) (len=6) {
app-1 | (string) (len=5) "depth": (float64) 1,
app-1 | @@ -76,3 +68,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -84,4 +76,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -100,3 +90,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -108,4 +98,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/blog.test
app-1 | Messages: crawler output for http://blog.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/spa.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":1021, "status_code":200, "type":"script", "url":"http://spa.test/assets/app.js"}, map[string]interface {}{"size_bytes":569, "status_code":200, "type":"style", "url":"http://spa.test/assets/app.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"SPA Test"}, "status":"ok", "url":"http://spa.test"}}, "root_url":"http://spa.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":1021, "status_code":200, "type":"script", "url":"http://spa.test/assets/app.js"}, map[string]interface {}{"size_bytes":569, "status_code":200, "type":"style", "url":"http://spa.test/assets/app.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"SPA Test"}, "status":"ok", "url":"http://spa.test"}}, "root_url":"http://spa.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=2) {
app-1 | @@ -18,4 +18,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/spa.test
app-1 | Messages: crawler output for http://spa.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/missing.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":interface {}(nil), "broken_links":interface {}(nil), "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":interface {}(nil), "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,5 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=9) {
app-1 | - (string) (len=6) "assets": (interface {}) <nil>,
app-1 | - (string) (len=12) "broken_links": (interface {}) <nil>,
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -11,9 +9,3 @@
app-1 | (string) (len=11) "http_status": (float64) 0,
app-1 | - (string) (len=3) "seo": (map[string]interface {}) (len=5) {
app-1 | - (string) (len=11) "description": (string) "",
app-1 | - (string) (len=15) "has_description": (bool) false,
app-1 | - (string) (len=6) "has_h1": (bool) false,
app-1 | - (string) (len=9) "has_title": (bool) false,
app-1 | - (string) (len=5) "title": (string) ""
app-1 | - },
app-1 | + (string) (len=3) "seo": (interface {}) <nil>,
app-1 | (string) (len=6) "status": (string) (len=5) "error",
app-1 | Test: TestCLIIntegrationAgainstFixtures/missing.test
app-1 | Messages: crawler output for http://missing.test must match fixture
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures (2.45s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/simple.test (1.76s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/example.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/blog.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/spa.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/missing.test (0.45s)
app-1 | FAIL
app-1 | FAIL project 2.452s
app-1 | FAIL
app-1 | make: \*\*\* [Makefile:35: test] Error 1

app-1 exited with code 2
Aborting on container exit...
Container source-app-1 Stopping
Container source-app-1 Stopped
Container source-server-1 Stopping
Container source-server-1 Stopped

Error: The tests have failed. Examine what they have to say. Inhale deeply. Exhale. Fix the code.
Error: The process '/usr/bin/docker' failed with exit code 2
at ExecState.\_setResult (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:206396)
at ExecState.CheckComplete (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:205956)
at ChildProcess.<anonymous> (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:204797)
1s
0s

time="2026-02-26T14:44:19Z" level=warning msg="Found orphan containers ([source-app-run-56086eaab9b0]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up."
Container source-server-1 Running
Container source-app-1 Creating
Container source-app-1 Created
Attaching to app-1, server-1
app-1 | go test -v ./tests
app-1 | === RUN TestAnalyzeJSONReportSingleDepth
app-1 | analyze_test.go:141: expected 1 page in report, got 3
app-1 | --- FAIL: TestAnalyzeJSONReportSingleDepth (0.00s)
app-1 | === RUN TestCLIIntegrationAgainstFixtures
app-1 | === RUN TestCLIIntegrationAgainstFixtures/simple.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Simple Test Site"}, "status":"ok", "url":"http://simple.test"}}, "root_url":"http://simple.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Simple Test Site"}, "status":"ok", "url":"http://simple.test"}}, "root_url":"http://simple.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,7 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | - (string) (len=6) "assets": ([]interface {}) {
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | + (map[string]interface {}) (len=6) {
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | Test: TestCLIIntegrationAgainstFixtures/simple.test
app-1 | Messages: crawler output for http://simple.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/example.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":234, "status_code":200, "type":"image", "url":"http://example.test/assets/img/hero.svg"}, map[string]interface {}{"size_bytes":194, "status_code":200, "type":"script", "url":"http://example.test/assets/site.js"}, map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test"}, "status":"ok", "url":"http://example.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"About Example Test"}, "status":"ok", "url":"http://example.test/about.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test News"}, "status":"ok", "url":"http://example.test/news.html"}}, "root_url":"http://example.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":234, "status_code":200, "type":"image", "url":"http://example.test/assets/img/hero.svg"}, map[string]interface {}{"size_bytes":194, "status_code":200, "type":"script", "url":"http://example.test/assets/site.js"}, map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test"}, "status":"ok", "url":"http://example.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"About Example Test"}, "status":"ok", "url":"http://example.test/about.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":242, "status_code":200, "type":"style", "url":"http://example.test/assets/styles.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Example Test News"}, "status":"ok", "url":"http://example.test/news.html"}}, "root_url":"http://example.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=3) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=3) {
app-1 | @@ -25,4 +25,2 @@
app-1 | },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -40,3 +38,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -48,4 +46,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -64,3 +60,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -72,4 +68,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/example.test
app-1 | Messages: crawler output for http://example.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/blog.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"assets":[]interface {}{}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "broken_links":[]interface {}{}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":240, "status_code":200, "type":"image", "url":"http://blog.test/assets/img/cover.svg"}, map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Blog Test"}, "status":"ok", "url":"http://blog.test"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Архив"}, "status":"ok", "url":"http://blog.test/archive.html"}, map[string]interface {}{"depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"Crawler Blog"}, "status":"ok", "url":"http://blog.test/feed.xml"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Робот"}, "status":"ok", "url":"http://blog.test/posts/post-1.html"}, map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":226, "status_code":200, "type":"style", "url":"http://blog.test/assets/style.css"}}, "depth":1, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":true, "has_title":true, "title":"Sitemap"}, "status":"ok", "url":"http://blog.test/posts/post-2.html"}}, "root_url":"http://blog.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=5) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=2) {
app-1 | @@ -19,4 +19,2 @@
app-1 | },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -34,3 +32,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -42,4 +40,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -58,7 +54,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | - (string) (len=6) "assets": ([]interface {}) {
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | - },
app-1 | + (map[string]interface {}) (len=6) {
app-1 | (string) (len=5) "depth": (float64) 1,
app-1 | @@ -76,3 +68,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -84,4 +76,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | @@ -100,3 +90,3 @@
app-1 | },
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=1) {
app-1 | @@ -108,4 +98,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/blog.test
app-1 | Messages: crawler output for http://blog.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/spa.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":1021, "status_code":200, "type":"script", "url":"http://spa.test/assets/app.js"}, map[string]interface {}{"size_bytes":569, "status_code":200, "type":"style", "url":"http://spa.test/assets/app.css"}}, "broken_links":[]interface {}{}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"SPA Test"}, "status":"ok", "url":"http://spa.test"}}, "root_url":"http://spa.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":[]interface {}{map[string]interface {}{"size_bytes":1021, "status_code":200, "type":"script", "url":"http://spa.test/assets/app.js"}, map[string]interface {}{"size_bytes":569, "status_code":200, "type":"style", "url":"http://spa.test/assets/app.css"}}, "depth":0, "discovered_at":"", "http_status":200, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":true, "title":"SPA Test"}, "status":"ok", "url":"http://spa.test"}}, "root_url":"http://spa.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,3 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=8) {
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=6) "assets": ([]interface {}) (len=2) {
app-1 | @@ -18,4 +18,2 @@
app-1 | }
app-1 | - },
app-1 | - (string) (len=12) "broken_links": ([]interface {}) {
app-1 | },
app-1 | Test: TestCLIIntegrationAgainstFixtures/spa.test
app-1 | Messages: crawler output for http://spa.test must match fixture
app-1 | === RUN TestCLIIntegrationAgainstFixtures/missing.test
app-1 | cli_test.go:47:
app-1 | Error Trace: /project/tests/cli_test.go:102
app-1 | /project/tests/cli_test.go:47
app-1 | Error: Not equal:
app-1 | expected: map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"assets":interface {}(nil), "broken_links":interface {}(nil), "depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":map[string]interface {}{"description":"", "has_description":false, "has_h1":false, "has_title":false, "title":""}, "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 | actual : map[string]interface {}{"depth":10, "generated_at":"", "pages":[]interface {}{map[string]interface {}{"depth":0, "discovered_at":"", "error":"Get \"http://missing.test\": dial tcp: lookup missing.test on 127.0.0.11:53: no such host", "http_status":0, "seo":interface {}(nil), "status":"error", "url":"http://missing.test"}}, "root_url":"http://missing.test"}
app-1 |
app-1 | Diff:
app-1 | --- Expected
app-1 | +++ Actual
app-1 | @@ -4,5 +4,3 @@
app-1 | (string) (len=5) "pages": ([]interface {}) (len=1) {
app-1 | - (map[string]interface {}) (len=9) {
app-1 | - (string) (len=6) "assets": (interface {}) <nil>,
app-1 | - (string) (len=12) "broken_links": (interface {}) <nil>,
app-1 | + (map[string]interface {}) (len=7) {
app-1 | (string) (len=5) "depth": (float64) 0,
app-1 | @@ -11,9 +9,3 @@
app-1 | (string) (len=11) "http_status": (float64) 0,
app-1 | - (string) (len=3) "seo": (map[string]interface {}) (len=5) {
app-1 | - (string) (len=11) "description": (string) "",
app-1 | - (string) (len=15) "has_description": (bool) false,
app-1 | - (string) (len=6) "has_h1": (bool) false,
app-1 | - (string) (len=9) "has_title": (bool) false,
app-1 | - (string) (len=5) "title": (string) ""
app-1 | - },
app-1 | + (string) (len=3) "seo": (interface {}) <nil>,
app-1 | (string) (len=6) "status": (string) (len=5) "error",
app-1 | Test: TestCLIIntegrationAgainstFixtures/missing.test
app-1 | Messages: crawler output for http://missing.test must match fixture
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures (2.05s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/simple.test (1.62s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/example.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/blog.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/spa.test (0.08s)
app-1 | --- FAIL: TestCLIIntegrationAgainstFixtures/missing.test (0.20s)
app-1 | FAIL
app-1 | FAIL project 2.056s
app-1 | FAIL
app-1 | make: \*\*\* [Makefile:35: test] Error 1

app-1 exited with code 2
Aborting on container exit...
Container source-app-1 Stopping
Container source-app-1 Stopped
Container source-server-1 Stopping
Container source-server-1 Stopped

Error: The tests have failed. Examine what they have to say. Inhale deeply. Exhale. Fix the code.
Error: The process '/usr/bin/docker' failed with exit code 2
at ExecState.\_setResult (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:206396)
at ExecState.CheckComplete (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:205956)
at ChildProcess.<anonymous> (file:///home/runner/work/\_actions/hexlet/project-action/release/dist/run-tests/index.js:2:204797)
1s
0s
