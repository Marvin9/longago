## API Endpoints

- [start](#Start)
- [pause](#Pause)
- [resume](#Resume)
- [stop](#Stop)

# Start

**URL** : ```/p/start```

**METHOD** : ```POST```

**DATA** : File which user wants to upload on server.

**Example**

``` curl -F file=@"./fixtures/100000.csv" localhost:8000/p/start ```

OR

```make start-request```

#### Success Response

**Code** ```200 OK```

**JSON**

```json
{
    "error": false,
    "data": "[timestamp]__[filename]"
}
// use "data" as instance_id to pause/resume/stop upload in next request(s)
```
#### Error Response

**Code** ```409 Status Conflict``` OR ```500 Internal Server Error```

**JSON**

```json
{
    "error": true,
    "data": "Reason why request failed."
}
```

# Pause

**URL** : ```/p/pause```

**METHOD** : ```POST```

**DATA** :
```
{
    "instance_id": "id that returned on start request."
}
```

**Example**

``` curl -H "content-type: application/json" -X POST -d '{"instance_id": "instance_id"}' localhost:8000/p/pause ```

OR

```make instance_id=i pause-request```

#### Success Response

**Code** ```200 OK```

**JSON**

```json
{
    "error": false,
    "data": "instance_id"
}
```

#### Error Response

**Code** ```400 Bad Request```

**JSON**

```json
{
    "error": true,
    "data": "Reason why request failed."
}
```

# Resume

**URL** : ```/p/resume```

**METHOD** : ```POST```

**DATA** : File which was paused, along with instance_id in form-data.

**Example**

``` curl -F file=@"./fixtures/100000.csv" -F instance_id="instance_id" localhost:8000/p/resume  ```

OR

```make instance_id=i resume-request```

#### Success Response

**Code** ```200 OK```

**JSON**

```json
{
    "error": false,
    "data": "instance_id"
}
```

#### Error Response

**Code** ```400 Bad Request``` OR ```409 Conflict``` OR ```500 Internal Server Error```

**JSON**

```json
{
    "error": true,
    "data": "Reason why request failed."
}
```

# Stop

**URL** : ```/p/stop```

**METHOD** : ```POST```

**DATA** :

```
{
    "instance_id": "id that returned on start request."
}
```

**Example**

``` curl -H "content-type: application/json" -X POST -d '{"instance_id": "instance_id"}' localhost:8000/p/stop  ```

OR

```make instance_id=i stop-request```

#### Success Response

**Code** ```200 OK```

**JSON**

```json
{
    "error": false,
    "data": "Success message."
}
```

#### Error Response

**Code** ```400 Bad Request``` OR ```409 Conflict```

**JSON**

```json
{
    "error": true,
    "data": "Reason why request failed."
}
```