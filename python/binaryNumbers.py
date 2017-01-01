#!/usr/bin/env python3

from sys import stdin

def get_largest( num, lar_bin=1, times=0 ):

	times += 1
	nlar_bin = lar_bin * 2
	if nlar_bin > num:
		return [ lar_bin, times ]
	else:
		return get_largest( num, nlar_bin, times )

def get_binstruct( num ):

	nl = get_largest( num )
	if ( num != 0 ):
		lnum = nl[0]
		times = nl[1]
		return nl + get_binstruct( num - lnum )
	else:
		return nl[:-2]

def clean( ar ):

	narr = []
	for i in range( 1, len( ar ), 2 ):
		narr.append( ar[i] )

	return narr

def get_binstr( binmap ):
	
	binstr = ""
	largest = binmap[0]
	for i in range( 1, largest + 1 ):
	
		if i in binmap:
			binstr = "1" + binstr
		else:
			binstr = "0" + binstr

	return binstr

def get_largest_binstr_len( binstr ):
	binstrings = binstr.split( "0" )

	lngst = 0;
	for s in binstrings:
		strlen = len( s )
		if strlen > lngst:
			lngst = strlen

	return lngst


num = int( stdin.readline() )

if num != 0:

	binmap = clean( get_binstruct( num ) )
	binstr = get_binstr( binmap )
	largest_length = get_largest_binstr_len( binstr )

	print( largest_length )

else:
	print( 0 )

