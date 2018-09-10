<html>
<head>
<title></title>
</head>
<body>
<input type="checkbox" name="interest" value="football">足球</input>
<input type="checkbox" name="interest" value="basketball">篮球</input>
<input type="checkbox" name="interest" value="tennise">网球</input>
    
Welcome, {{.UserName}}
<form action='/count' method='get'>
<input type="submit" value="count"/>
</form>
<form enctype="multipart/form-data" action="/upload" method="post">
<input type="file" name="uploadfile" />
<input type="hidden" name="token" value="{{.}}"/>
<input type="submit" value="upload"/>
</form>
</body>
</html>