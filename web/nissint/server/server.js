#!/usr/bin/env nodejs

var express = require('express')
var app = express()

app.use( express.static( 'client' ) );

app.listen( 6699, () => {
  console.log('Example app listening on port 3000!')
})

