#!/usr/bin/env nodejs

const https = require( 'https' )
const fs = require( 'fs' )

const opts = {
	method: "GET",
	hostname: "www.google.com",
	port: 443,
	path: "/"
}

var req = https.request( opts, ( res ) => {
	
	res.setEncoding( 'utf8' );
	res.on( "data", ( d ) => {
		console.log( d );
	})

})


req.end();

console.log( "HI!" )

