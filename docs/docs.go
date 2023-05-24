// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/permission/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新权限",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "更新权限",
                "parameters": [
                    {
                        "description": "权限信息",
                        "name": "perm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PermissionModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-entity_PermissionEntity"
                        }
                    }
                }
            }
        },
        "/permission/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "权限管理列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限管理列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "权限动作：比如get、post、delete等",
                        "name": "action",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序类型：asc、desc",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "权限名称",
                        "name": "permName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "权限类型：1-菜单、2-按钮",
                        "name": "permType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "URL路径",
                        "name": "url",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-model_PageData-model_PermissionModel"
                        }
                    }
                }
            }
        },
        "/role/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "添加角色",
                "parameters": [
                    {
                        "description": "角色信息",
                        "name": "role",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.RoleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-entity_RoleEntity"
                        }
                    }
                }
            }
        },
        "/role/assignPermission": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分配权限",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "分配权限",
                "parameters": [
                    {
                        "description": "角色与权限信息",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RoleAssignPermModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-any"
                        }
                    }
                }
            }
        },
        "/role/delete/:id": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-any"
                        }
                    }
                }
            }
        },
        "/role/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色管理列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色管理列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序类型：asc、desc",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色编号",
                        "name": "roleCode",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "roleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-model_PageData-model_RoleModel"
                        }
                    }
                }
            }
        },
        "/test/query": {
            "get": {
                "description": "get list by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "test data list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序类型：asc、desc",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-model_PageData-model_TestResult"
                        }
                    }
                }
            }
        },
        "/user/assignRole": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "给用户分配角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "给用户分配角色",
                "parameters": [
                    {
                        "description": "用户ID与角色编号",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AssignRoleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-any"
                        }
                    }
                }
            }
        },
        "/user/createByEmail": {
            "post": {
                "description": "邮箱注册用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "邮箱注册用户",
                "parameters": [
                    {
                        "description": "邮箱与密码",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserEmailRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-model_UserModel"
                        }
                    }
                }
            }
        },
        "/user/delete/:id": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "刪除用戶",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "刪除用戶",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-any"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户管理列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户管理列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序类型：asc、desc",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "realName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "状态: 1-启用，2-禁用",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-model_PageData-model_UserModel"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "使用邮箱登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "使用邮箱登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.PermissionEntity": {
            "type": "object",
            "properties": {
                "action": {
                    "description": "权限动作：比如get、post、delete等",
                    "type": "string"
                },
                "createAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "createBy": {
                    "description": "创建人",
                    "type": "integer"
                },
                "deleted": {
                    "description": "是否删除 1-否，2-是",
                    "type": "integer"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父级ID：资源层级关系（0表示顶级）",
                    "type": "integer"
                },
                "permName": {
                    "description": "权限名称",
                    "type": "string"
                },
                "permType": {
                    "description": "权限类型：1-菜单、2-按钮",
                    "type": "integer"
                },
                "updateAt": {
                    "description": "更新时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "updateBy": {
                    "description": "更新人",
                    "type": "integer"
                },
                "url": {
                    "description": "URL路径",
                    "type": "string"
                }
            }
        },
        "entity.RoleEntity": {
            "type": "object",
            "properties": {
                "createAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "createBy": {
                    "description": "创建人",
                    "type": "integer"
                },
                "deleted": {
                    "description": "是否删除 1-否，2-是",
                    "type": "integer"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "roleCode": {
                    "description": "角色编码",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "updateAt": {
                    "description": "更新时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "updateBy": {
                    "description": "更新人",
                    "type": "integer"
                }
            }
        },
        "localtime.LocalTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "model.AssignRoleModel": {
            "type": "object",
            "required": [
                "roleCode",
                "userId"
            ],
            "properties": {
                "roleCode": {
                    "description": "角色编码",
                    "type": "string"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "model.PageData-model_PermissionModel": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "总记录数",
                    "type": "integer"
                },
                "data": {
                    "description": "分页数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PermissionModel"
                    }
                },
                "pageNo": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "model.PageData-model_RoleModel": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "总记录数",
                    "type": "integer"
                },
                "data": {
                    "description": "分页数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RoleModel"
                    }
                },
                "pageNo": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "model.PageData-model_TestResult": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "总记录数",
                    "type": "integer"
                },
                "data": {
                    "description": "分页数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TestResult"
                    }
                },
                "pageNo": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "model.PageData-model_UserModel": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "总记录数",
                    "type": "integer"
                },
                "data": {
                    "description": "分页数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserModel"
                    }
                },
                "pageNo": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "model.PermissionModel": {
            "type": "object",
            "required": [
                "action",
                "permName",
                "permType",
                "url"
            ],
            "properties": {
                "action": {
                    "description": "权限动作：比如get、post、delete等",
                    "type": "string"
                },
                "id": {
                    "description": "权限ID",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父级ID：资源层级关系",
                    "type": "integer"
                },
                "permName": {
                    "description": "权限名称",
                    "type": "string"
                },
                "permType": {
                    "description": "权限类型：1-菜单、2-按钮",
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 1
                },
                "url": {
                    "description": "URL路径",
                    "type": "string"
                }
            }
        },
        "model.RoleAssignPermModel": {
            "type": "object",
            "required": [
                "permIdList",
                "roleId"
            ],
            "properties": {
                "permIdList": {
                    "description": "权限ID列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "roleId": {
                    "description": "角色ID",
                    "type": "integer"
                }
            }
        },
        "model.RoleModel": {
            "type": "object",
            "required": [
                "roleCode",
                "roleName"
            ],
            "properties": {
                "createAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "createBy": {
                    "description": "创建人",
                    "type": "integer"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "roleCode": {
                    "description": "角色编号",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "updateAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "updateBy": {
                    "description": "更新人",
                    "type": "integer"
                }
            }
        },
        "model.TestResult": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "model.UserEmailRegister": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "model.UserLoginModel": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "model.UserModel": {
            "type": "object",
            "properties": {
                "createAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "createBy": {
                    "description": "创建人",
                    "type": "integer"
                },
                "deleted": {
                    "description": "是否删除 1-否，2-是",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "realName": {
                    "description": "姓名",
                    "type": "string"
                },
                "roleCode": {
                    "description": "角色编号",
                    "type": "string"
                },
                "status": {
                    "description": "状态: 1-启用，2-禁用",
                    "type": "integer"
                },
                "updateAt": {
                    "description": "更新时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/localtime.LocalTime"
                        }
                    ]
                },
                "updateBy": {
                    "description": "更新人",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "result.Result-any": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-entity_PermissionEntity": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.PermissionEntity"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-entity_RoleEntity": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.RoleEntity"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-model_PageData-model_PermissionModel": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PageData-model_PermissionModel"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-model_PageData-model_RoleModel": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PageData-model_RoleModel"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-model_PageData-model_TestResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PageData-model_TestResult"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-model_PageData-model_UserModel": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PageData-model_UserModel"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "result.Result-model_UserModel": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码，200表示成功，其它为失败",
                    "type": "string"
                },
                "data": {
                    "description": "响应数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserModel"
                        }
                    ]
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "timeStamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8089",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "XXX API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
