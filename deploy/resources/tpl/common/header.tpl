{{define "header"}}
<!DOCTYPE html>
<html>   
	<head>        
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>        
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>        
	<meta name="description" content="{{"meta.description"|get}}"/>        
	<meta name="keywords" content="{{"meta.keywords"|get}}"/>        
	<meta name="author" content="{{"meta.author"|get}}"/>        
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
	<link rel="stylesheet" href="/prettify/normalize.css"/>
	<link rel="stylesheet" href="/prettify/prettify_{{"codetheme"|get}}.css"/>
	<link rel="stylesheet" href="/css/main.css"/>
	<link rel="shortcut icon" href="/fav.ico"/>
	<script type="text/javascript" src="/prettify/prettify.js"></script>
	</head>
	<body onload="prettyPrint()">
{{end}}