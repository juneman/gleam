package ui

import (
	"text/template"
)

var JobStatusTpl = template.Must(template.New("job").Parse(`<!DOCTYPE html>
<html>
  <head>
    <title>Job {{ .Status.Id }}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
  </head>
  <body>
    <div class="container">
      <div class="page-header">
	    <h1>
          <a href="https://github.com/chrislusf/gleam">Gleam</a> <small>{{ .Version }}</small>
	    </h1>
      </div>

      <div class="row">
        <div class="col-sm-6">
          {{ with .Status.Driver }}
          <h2>Driver Program</h2>
          <table class="table">
            <tbody>
              <tr>
                <th>User</th>
                <td>{{ .Username }}</td>
              </tr>
              <tr>
                <th>Host</th>
                <td>{{ .Hostname }}</td>
              </tr>
              <tr>
                <th>Executable</th>
                <td style="max-width:150px;word-wrap:break-word;">{{ .Executable }}</td>
              </tr>
              <tr>
                <th>Start</th>
                <td>{{ .StartTime }}</td>
              </tr>
              <tr>
                <th>Stop</th>
                <td>{{ .StopTime }}</td>
              </tr>
            </tbody>
          </table>
          {{ end }}
        </div>

        <div class="col-sm-6">
          <h2>System Stats</h2>
          <table class="table table-condensed table-striped">
            <tr>
              <th>Jobs Completed</th>
              <td>100</td>
            </tr>
          </table>
        </div>
      </div>

      <p>{{.Svg}}

      {{ with .Status.TaskGroups }}
      <div class="row">
        <h2>Task Group</h2>
        <table class="table table-striped">
          <thead>
            <tr>
              <th>Steps</th>
              <th>Name</th>
              <th>IO</th>
              <th>Allocation</th>
              <th>CPU</th>
              <th>Memory</th>
            </tr>
          </thead>
          <tbody>
          {{ range $tg_index, $tg := . }}
            <tr>
              <td>{{ $tg.StepIds }}</td>
              <td>{{with $tg.Request}}{{.Name}}{{end}}</td>
              <td>
                {{with $tg.Request}}{{with .Instructions}}
                {{ range $inst_index, $inst := .Instructions }}
                    {{with .InputShardLocations}}Input: <ul>{{ range . }}<li>{{.Name}}@{{.Host}}:{{.Port}}</li>{{end}}</ul>{{end}}
                    {{with .OutputShardLocations}}Output:<ul>{{ range . }}<li>{{.Name}}@{{.Host}}:{{.Port}}</li>{{end}}{{end}}
                {{end}}
                {{ end }}{{ end }}
              </td>
              <td>{{with $tg.Allocation}}
                    {{.Location.DataCenter}}-{{.Location.Rack}}-{{.Location.Server}}:{{.Location.Port}}
                    <br/>
                    CPU:{{.Allocated.CpuCount}} Memory:{{.Allocated.MemoryMb}}MB
                  {{end}}</td>
              <td><ul>{{range .Executions}}
                   <li>
                     Start: {{.StartTime}}<br/>
                     Stop:  {{.StopTime}}<br/>
                     System:{{.SystemTime}} Seconds<br/>
                     User:{{.UserTime}} Seconds
                   </li>
                  {{end}}</ul></td>
              <td>{{with $tg.Request}}{{.Resource.MemoryMb}}{{end}}</td>
            </tr>
          {{ end }}
          </tbody>
        </table>
      </div>
      {{ end }}

    </div>
  </body>
</html>
`))
