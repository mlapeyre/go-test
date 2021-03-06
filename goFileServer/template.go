package goFileServer

func getListTemplate() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title></title>
    <meta name="description" content="">
    <link href="http://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<div class="container">
    <div class="row">
        <table class="table table-condensed table-hover table-striped">
            <thead>
            <tr>
                <th>Name</th>
                <th>Size</th>
                <th>Rights</th>
                <th>Modification date</th>
            </tr>
            </thead>
            <tbody>
            {{range .}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{if .IsDir}} - {{else}} {{.Size}} {{end}}</td>
                <td>{{.Mode}}</td>
                <td>{{.ModTime }}</td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>
</body>
`
}

