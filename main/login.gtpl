<html>
<head>
<title></title>
</head>
<body>
<form action='/login' method='post'>
<input type="checkbox" name="interest" value="football">足球</input>
<input type="checkbox" name="interest" value="basketball">篮球</input>
<input type="checkbox" name="interest" value="tennise">网球</input>
    User name : <input type="text" name="username" value="{{.username}}">
    Password: <input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="login">


</form>
<form enctype="multipart/form-data" action="/upload" method="post">
<input type="file" name="uploadfile" />
<input type="hidden" name="token" value="{{.}}"/>
<input type="submit" value="upload"/>
</form>
</body>
</html>