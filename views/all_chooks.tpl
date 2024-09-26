<!DOCTYPE html>

<html>
  <head>
    <title>Chookpedia</title>
  </head>

  <body>
    <h1>Chookpedia</h1>
    
    <table>
    {{range $key, $val := .chooks}}
      <tr>
	<td><a href="/chook/{{$val.Id}}">{{$val.Name}}</a></td>
	<td>{{$val.Breed.Name}}
      </tr>
    {{end}}
    </table>
  </body>
</html>
