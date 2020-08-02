#Account management system server for eXternOS

Endpoints:

`/api/newToken` Get new token by passing login and password in !JSON! body in GET request
`/api/register` You need to pass these parameters in JSON body request:
```golang 
        Name      string `json:"name"` (obligatory)
	Username  string `json:"username"` (obligatory)
	Website   string `json:"website"` 
	Email     string `json:"email"`
	Avatarurl string `json:"avatarurl"` (obligatory)
	Password  string `json:"password"` (obligatory)
```


To be updated