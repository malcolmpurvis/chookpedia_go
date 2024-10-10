<!DOCTYPE html>

<html>
  <head>
    <title>{{.Name}} the Chook</title>
  </head>

  <body>
    <h1>{{.Name}}</h1>

    <img src="{{.PhotoURL}}">

    <dl>
      <dt> Name: </dt> <dd> {{.Name}} </dd>
      <dt>Colour:</dt> <dd> {{.Colour}}</dd>
      <dt>Breed:</dt> <dd>{{.Breed}}</dd>
    </dl>
  </body>
</html>
