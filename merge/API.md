端口：90

## POST：

---

方法描述：用户注册

URL地址：/register

请求方法：POST

请求参数：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  username  |  string  |  用户名  |
|  password  |  string  |  密码  |

例如

```
{
    "username":"abc",
    "password":"abc"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |
|  data.ID  |  string  |  用户名  |
|  data.Token  |  string  |  基于 jwt 的令牌  |

返回示例

```
{
    "data": {
        "ID": "abc",
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImFiYyIsInBhc3N3b3JkIjoiYWJjIiwiZXhwIjoxNjA3Nzg1MjY1LCJpc3MiOiJseHkiLCJuYmYiOjE2MDc3ODA2NjV9.51RLPkF8olOM8sE2U7McuOpR1dI4k3o6m0Df2aHi8g0"
    },
    "message": "成功",
    "status": 200
}
```

---

方法描述：用户登录

URL地址：/login

请求方法：POST

请求参数：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  username  |  string  |  用户名  |
|  password  |  string  |  密码  |

例如

```
{
    "username":"abc",
    "password":"abc"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |
|  data.ID  |  string  |  用户名  |
|  data.Token  |  string  |  基于 jwt 的令牌  |

返回示例

```
{
    "data": {
        "ID": "abc",
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImFiYyIsInBhc3N3b3JkIjoiYWJjIiwiZXhwIjoxNjA3Nzg1MjY1LCJpc3MiOiJseHkiLCJuYmYiOjE2MDc3ODA2NjV9.51RLPkF8olOM8sE2U7McuOpR1dI4k3o6m0Df2aHi8g0"
    },
    "message": "登录成功",
    "status": 200
}
```

---

方法描述：新建建筑

URL地址：/newbuilding

请求方法：POST

请求参数：

请求头需携带token

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  building_name  |  string  |  建筑名  |
|  building_class  |  string  |  建筑类型  |
|  text_src  |  string  |  建筑简介  |

例如

```
{
    "building_name":"testbook",
    "building_class":"book",
    "text_src":"test comment"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |

返回示例

```
{
    "message": "成功",
    "status": 200
}
```

---

方法描述：新建文章

URL地址：/newarticle

请求方法：POST

请求参数：

请求头需携带token

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  name  |  string  |  所属建筑名  |
|  up_id  |  string  |  发布人用户名  |
|  text_src  |  string  |  文章内容  |

例如

```
{
	"name":"testbook",
    "up_id":"abc",
    "text_src":"abc's text."
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |

返回示例

```
{
    "message": "成功",
    "status": 200
}
```
---

方法描述：根据建筑名展示其附属文章

URL地址：/showallarticle

请求方法：POST

请求参数：

请求头需携带token

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  building_name  |  string  |  建筑名  |

例如

```
{
    "building_name":"testbook"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |
|  data[].name  |  string  |  文章所属建筑名  |
|  data[].up_id  |  string  |  文章发布人用户名  |
|  data[].up_time  |  string  |  文章发布时间  |
|  data[].text_src  |  string  |  文章内容  |

返回示例

```
{
    "data": [
        {
            "name": "testbook",
            "up_id": "abc",
            "up_time": 1607783394,
            "text_src": "abc's text2.",

        },
        {
            "name": "testbook",
            "up_id": "abc",
            "up_time": 1607783452,
            "text_src": "abc's text3.",
        }
    ]
    "message": "成功",
    "status": 200
}
```
---

方法描述：用户保存的文章（发布的文章）

URL地址：/querysave

请求方法：POST

请求参数：

请求头需携带token

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  username  |  string  |  用户名  |


例如

```
{
    "username":"abc"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |
|  data[].name  |  string  |  文章所属建筑名  |
|  data[].up_id  |  string  |  文章发布人用户名  |
|  data[].up_time  |  int  |  文章发布时间  |
|  data[].text_src  |  string  |  文章内容  |


返回示例

```
{
    "data": [
        {
            "name": "testbook",
            "up_id": "abc",
            "up_time": 1607783394,
            "text_src": "abc's text2.",

        },
        {
            "name": "testbook",
            "up_id": "abc",
            "up_time": 1607783452,
            "text_src": "abc's text3.",
        }
    ]
    "message": "请求成功",
    "status": 200
}
```
---

方法描述：用户保存的建筑（发布的文章所属的建筑）

URL地址：/querysavebuilding

请求方法：POST

请求参数：

请求头需携带token

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  username  |  string  |  用户名  |


例如

```
{
    "username":"abc"
}
```

返回参数
| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  message  |  string  |  返回信息  |
|  status  |  int  |  状态  |
|  data[].building_name  |  string  |  建筑名  |
|  data[].building_class  |  string  |  建筑类型  |
|  data[].text_src  |  string  |  建筑内容  |
|  data[].start_time  |  int  |  建筑发布时间  |


返回示例

```
{
    "data": [
        {
            "building_name": "testbook",
            "building_class": "book",
            "text_src": "test comment",
            "start_time": 1607782406,
        }
    ],
    "message": "请求成功",
    "status": 200
}
```

---
