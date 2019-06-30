package adl_test

import (
	"encoding/json"
	"testing"

	"github.com/wxio/tron-go/adl"
)

func TestAst2Go(t *testing.T) {
	m := adl.Module{}
	err := json.Unmarshal([]byte(ast_oneofeach), &m)
	if err != nil {
		t.Fatal(err)
	}

}

const adl_oneofeach = `
module helix.protoapp.requests {

	import common.http.*;
	import common.*;
	import common.db.DbTable;
	import sys.types.*;
	
	// comment
	/// doccmt
	@Path "localanno"
	type Hello<A> = Post<HelloReq, HelloResp<Vector<Vector<Vector<A>>>,Int32,Float>>;
	
	struct Const {
	  
	};
	
	type Literal<T> = StringMap<T>;
	type StrLiteral = Literal<String>;
	annotation StrLiteral { "a" : "b" } ;
	
	struct A {
	  String a;
	};
	struct B {
	  Int32 b;
	};
	union MyConfig {
	  A a;
	  B b;
	};
	type MyConfigMap = StringMap<MyConfig>;
	annotation MyConfig { "a" : { "a" : "mystring" } };
	annotation MyConfig { "b" : { "b" : 10 } };
	annotation MyConfigMap { 
	  "a b" : { "b" : { "b" : 10 } },
	  "a a" : { "a" : { "a" : "hw" } }
	};
	
	/// doconstr
	@SA { "a" : "b" }
	struct HelloReq {
	  String name;
	};
	
	struct SA {
	  String a;
	};
	
	struct HelloResp<A,B,C> {
	};
	
	/// docontype
	@Path "/debug/time"
	type CurrentTime = Get<Instant>;
	
	/// doconnewtype
	@Path "/debug/dummy-exception"
	newtype DummyException = Post<String, Unit>;
	
	union LoginResp<T> {
	  @SA { "a" : "b" }
	  T accessToken;
	};
	
	annotation Path "mod anno str";
	
	annotation HelloReq::name Path "field anno";
	
	annotation HelloReq DbTable {
	  "withIdPrimaryKey" : true,
	  "indexes" : [["username"]]
	};
	
	};	
`

// generated via timbod7's `adlc ast <file>`
const ast_oneofeach = `
{
    "annotations": [
        {
            "v1": {
                "moduleName": "common.http",
                "name": "Path"
            },
            "v2": "mod anno str"
        }
    ],
    "decls": {
        "CurrentTime": {
            "annotations": [
                {
                    "v1": {
                        "moduleName": "common.http",
                        "name": "Path"
                    },
                    "v2": "/debug/time"
                },
                {
                    "v1": {
                        "moduleName": "sys.annotations",
                        "name": "Doc"
                    },
                    "v2": "docontype\n"
                }
            ],
            "name": "CurrentTime",
            "type_": {
                "type_": {
                    "typeExpr": {
                        "parameters": [
                            {
                                "parameters": [],
                                "typeRef": {
                                    "reference": {
                                        "moduleName": "common",
                                        "name": "Instant"
                                    }
                                }
                            }
                        ],
                        "typeRef": {
                            "reference": {
                                "moduleName": "common.http",
                                "name": "Get"
                            }
                        }
                    },
                    "typeParams": []
                }
            },
            "version": "nothing"
        },
        "DummyException": {
            "annotations": [
                {
                    "v1": {
                        "moduleName": "common.http",
                        "name": "Path"
                    },
                    "v2": "/debug/dummy-exception"
                },
                {
                    "v1": {
                        "moduleName": "sys.annotations",
                        "name": "Doc"
                    },
                    "v2": "doconnewtype\n"
                }
            ],
            "name": "DummyException",
            "type_": {
                "newtype_": {
                    "default": "nothing",
                    "typeExpr": {
                        "parameters": [
                            {
                                "parameters": [],
                                "typeRef": {
                                    "primitive": "String"
                                }
                            },
                            {
                                "parameters": [],
                                "typeRef": {
                                    "reference": {
                                        "moduleName": "common",
                                        "name": "Unit"
                                    }
                                }
                            }
                        ],
                        "typeRef": {
                            "reference": {
                                "moduleName": "common.http",
                                "name": "Post"
                            }
                        }
                    },
                    "typeParams": []
                }
            },
            "version": "nothing"
        },
        "Hello": {
            "annotations": [
                {
                    "v1": {
                        "moduleName": "common.http",
                        "name": "Path"
                    },
                    "v2": "localanno"
                },
                {
                    "v1": {
                        "moduleName": "sys.annotations",
                        "name": "Doc"
                    },
                    "v2": "doccmt\n"
                }
            ],
            "name": "Hello",
            "type_": {
                "type_": {
                    "typeExpr": {
                        "parameters": [
                            {
                                "parameters": [],
                                "typeRef": {
                                    "reference": {
                                        "moduleName": "",
                                        "name": "HelloReq"
                                    }
                                }
                            },
                            {
                                "parameters": [
                                    {
                                        "parameters": [
                                            {
                                                "parameters": [
                                                    {
                                                        "parameters": [
                                                            {
                                                                "parameters": [],
                                                                "typeRef": {
                                                                    "typeParam": "A"
                                                                }
                                                            }
                                                        ],
                                                        "typeRef": {
                                                            "primitive": "Vector"
                                                        }
                                                    }
                                                ],
                                                "typeRef": {
                                                    "primitive": "Vector"
                                                }
                                            }
                                        ],
                                        "typeRef": {
                                            "primitive": "Vector"
                                        }
                                    },
                                    {
                                        "parameters": [],
                                        "typeRef": {
                                            "primitive": "Int32"
                                        }
                                    },
                                    {
                                        "parameters": [],
                                        "typeRef": {
                                            "primitive": "Float"
                                        }
                                    }
                                ],
                                "typeRef": {
                                    "reference": {
                                        "moduleName": "",
                                        "name": "HelloResp"
                                    }
                                }
                            }
                        ],
                        "typeRef": {
                            "reference": {
                                "moduleName": "common.http",
                                "name": "Post"
                            }
                        }
                    },
                    "typeParams": [
                        "A"
                    ]
                }
            },
            "version": "nothing"
        },
        "HelloReq": {
            "annotations": [
                {
                    "v1": {
                        "moduleName": "",
                        "name": "SA"
                    },
                    "v2": {
                        "a": "b"
                    }
                },
                {
                    "v1": {
                        "moduleName": "common.db",
                        "name": "DbTable"
                    },
                    "v2": {
                        "indexes": [
                            [
                                "username"
                            ]
                        ],
                        "withIdPrimaryKey": true
                    }
                },
                {
                    "v1": {
                        "moduleName": "sys.annotations",
                        "name": "Doc"
                    },
                    "v2": "doconstr\n"
                }
            ],
            "name": "HelloReq",
            "type_": {
                "struct_": {
                    "fields": [
                        {
                            "annotations": [
                                {
                                    "v1": {
                                        "moduleName": "common.http",
                                        "name": "Path"
                                    },
                                    "v2": "field anno"
                                }
                            ],
                            "default": "nothing",
                            "name": "name",
                            "serializedName": "name",
                            "typeExpr": {
                                "parameters": [],
                                "typeRef": {
                                    "primitive": "String"
                                }
                            }
                        }
                    ],
                    "typeParams": []
                }
            },
            "version": "nothing"
        },
        "HelloResp": {
            "annotations": [],
            "name": "HelloResp",
            "type_": {
                "struct_": {
                    "fields": [],
                    "typeParams": [
                        "A",
                        "B",
                        "C"
                    ]
                }
            },
            "version": "nothing"
        },
        "LoginResp": {
            "annotations": [],
            "name": "LoginResp",
            "type_": {
                "union_": {
                    "fields": [
                        {
                            "annotations": [
                                {
                                    "v1": {
                                        "moduleName": "",
                                        "name": "SA"
                                    },
                                    "v2": {
                                        "a": "b"
                                    }
                                }
                            ],
                            "default": "nothing",
                            "name": "accessToken",
                            "serializedName": "accessToken",
                            "typeExpr": {
                                "parameters": [],
                                "typeRef": {
                                    "typeParam": "T"
                                }
                            }
                        }
                    ],
                    "typeParams": [
                        "T"
                    ]
                }
            },
            "version": "nothing"
        },
        "SA": {
            "annotations": [],
            "name": "SA",
            "type_": {
                "struct_": {
                    "fields": [
                        {
                            "annotations": [],
                            "default": "nothing",
                            "name": "a",
                            "serializedName": "a",
                            "typeExpr": {
                                "parameters": [],
                                "typeRef": {
                                    "primitive": "String"
                                }
                            }
                        }
                    ],
                    "typeParams": []
                }
            },
            "version": "nothing"
        }
    },
    "imports": [
        {
            "moduleName": "sys.annotations"
        },
        {
            "moduleName": "common.http"
        },
        {
            "moduleName": "common"
        },
        {
            "scopedName": {
                "moduleName": "common.db",
                "name": "DbTable"
            }
        }
    ],
    "name": "helix.protoapp.requests"
}
`
