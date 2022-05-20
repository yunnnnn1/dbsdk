package myast

/*
{
  "SQLBigResult": false,
  "SQLBufferResult": false,
  "SQLCache": true,
  "SQLSmallResult": false,
  "CalcFoundRows": false,
  "StraightJoin": false,
  "Priority": 0,
  "ExplicitAll": false,
  "Distinct": false,
  "From": {
    "TableRefs": {
      "Left": {
        "Source": {
          "Schema": {
            "O": "test",
            "L": "test"
          },
          "Name": {
            "O": "t1",
            "L": "t1"
          },
          "DBInfo": null,
          "TableInfo": null,
          "IndexHints": [],
          "PartitionNames": [],
          "TableSample": null,
          "AsOf": null
        },
        "AsName": {
          "O": "",
          "L": ""
        }
      },
      "Right": null,
      "Tp": 0,
      "On": null,
      "Using": null,
      "NaturalJoin": false,
      "StraightJoin": false,
      "ExplicitParens": false
    }
  },
  "Where": {
    "Type": {
      "Tp": 0,
      "Flag": 0,
      "Flen": 0,
      "Decimal": 0,
      "Charset": "",
      "Collate": "",
      "Elems": null
    },
    "Op": 7,
    "L": {
      "Type": {
        "Tp": 0,
        "Flag": 0,
        "Flen": 0,
        "Decimal": 0,
        "Charset": "",
        "Collate": "",
        "Elems": null
      },
      "Name": {
        "Schema": {
          "O": "",
          "L": ""
        },
        "Table": {
          "O": "",
          "L": ""
        },
        "Name": {
          "O": "b",
          "L": "b"
        }
      },
      "Refer": null
    },
    "R": {
      "Type": {
        "Tp": 8,
        "Flag": 128,
        "Flen": 1,
        "Decimal": 0,
        "Charset": "binary",
        "Collate": "binary",
        "Elems": null
      }
    }
  },
  "Fields": {
    "Fields": [
      {
        "Offset": 7,
        "WildCard": null,
        "Expr": {
          "Type": {
            "Tp": 0,
            "Flag": 0,
            "Flen": 0,
            "Decimal": 0,
            "Charset": "",
            "Collate": "",
            "Elems": null
          },
          "Name": {
            "Schema": {
              "O": "",
              "L": ""
            },
            "Table": {
              "O": "",
              "L": ""
            },
            "Name": {
              "O": "a",
              "L": "a"
            }
          },
          "Refer": null
        },
        "AsName": {
          "O": "",
          "L": ""
        },
        "Auxiliary": false
      },
      {
        "Offset": 9,
        "WildCard": null,
        "Expr": {
          "Type": {
            "Tp": 0,
            "Flag": 0,
            "Flen": 0,
            "Decimal": 0,
            "Charset": "",
            "Collate": "",
            "Elems": null
          },
          "Name": {
            "Schema": {
              "O": "",
              "L": ""
            },
            "Table": {
              "O": "",
              "L": ""
            },
            "Name": {
              "O": "b",
              "L": "b"
            }
          },
          "Refer": null
        },
        "AsName": {
          "O": "",
          "L": ""
        },
        "Auxiliary": false
      }
    ]
  },
  "GroupBy": {
    "Items": [
      {
        "Expr": {
          "Type": {
            "Tp": 0,
            "Flag": 0,
            "Flen": 0,
            "Decimal": 0,
            "Charset": "",
            "Collate": "",
            "Elems": null
          },
          "Name": {
            "Schema": {
              "O": "",
              "L": ""
            },
            "Table": {
              "O": "",
              "L": ""
            },
            "Name": {
              "O": "a",
              "L": "a"
            }
          },
          "Refer": null
        },
        "Desc": false,
        "NullOrder": true
      },
      {
        "Expr": {
          "Type": {
            "Tp": 0,
            "Flag": 0,
            "Flen": 0,
            "Decimal": 0,
            "Charset": "",
            "Collate": "",
            "Elems": null
          },
          "Name": {
            "Schema": {
              "O": "",
              "L": ""
            },
            "Table": {
              "O": "",
              "L": ""
            },
            "Name": {
              "O": "b",
              "L": "b"
            }
          },
          "Refer": null
        },
        "Desc": false,
        "NullOrder": true
      }
    ]
  },
  "Having": null,
  "WindowSpecs": null,
  "OrderBy": {
    "Items": [
      {
        "Expr": {
          "Type": {
            "Tp": 0,
            "Flag": 0,
            "Flen": 0,
            "Decimal": 0,
            "Charset": "",
            "Collate": "",
            "Elems": null
          },
          "N": 1,
          "P": null,
          "Refer": null
        },
        "Desc": false,
        "NullOrder": true
      }
    ],
    "ForUnion": false
  },
  "Limit": {
    "Count": {
      "Type": {
        "Tp": 8,
        "Flag": 160,
        "Flen": 1,
        "Decimal": 0,
        "Charset": "binary",
        "Collate": "binary",
        "Elems": null
      }
    },
    "Offset": {
      "Type": {
        "Tp": 8,
        "Flag": 160,
        "Flen": 1,
        "Decimal": 0,
        "Charset": "binary",
        "Collate": "binary",
        "Elems": null
      }
    }
  },
  "LockInfo": null,
  "TableHints": null,
  "IsInBraces": false,
  "WithBeforeBraces": false,
  "QueryBlockOffset": 0,
  "SelectIntoOpt": null,
  "AfterSetOperator": null,
  "Kind": 0,
  "Lists": null,
  "With": null
}
*/

// SelectStmt  查询操作结构体
type SelectStmt struct {
	SQLBigResult    bool `json:"SQLBigResult"`
	SQLBufferResult bool `json:"SQLBufferResult"`
	SQLCache        bool `json:"SQLCache"`
	SQLSmallResult  bool `json:"SQLSmallResult"`
	CalcFoundRows   bool `json:"CalcFoundRows"`
	StraightJoin    bool `json:"StraightJoin"`
	Priority        int  `json:"Priority"`
	ExplicitAll     bool `json:"ExplicitAll"`
	Distinct        bool `json:"Distinct"`
	From            struct {
		TableRefs struct {
			Left struct {
				Source struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
					DBInfo         interface{}   `json:"DBInfo"`
					TableInfo      interface{}   `json:"TableInfo"`
					IndexHints     []interface{} `json:"IndexHints"`
					PartitionNames []interface{} `json:"PartitionNames"`
					TableSample    interface{}   `json:"TableSample"`
					AsOf           interface{}   `json:"AsOf"`
				} `json:"Source"`
				AsName struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"AsName"`
			} `json:"Left"`
			Right          interface{} `json:"Right"`
			Tp             int         `json:"Tp"`
			On             interface{} `json:"On"`
			Using          interface{} `json:"Using"`
			NaturalJoin    bool        `json:"NaturalJoin"`
			StraightJoin   bool        `json:"StraightJoin"`
			ExplicitParens bool        `json:"ExplicitParens"`
		} `json:"TableRefs"`
	} `json:"From"`
	Where struct {
		Type struct {
			Tp      int         `json:"Tp"`
			Flag    int         `json:"Flag"`
			Flen    int         `json:"Flen"`
			Decimal int         `json:"Decimal"`
			Charset string      `json:"Charset"`
			Collate string      `json:"Collate"`
			Elems   interface{} `json:"Elems"`
		} `json:"Type"`
		Op int `json:"Op"`
		L  struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
			Name struct {
				Schema struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Schema"`
				Table struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Table"`
				Name struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Name"`
			} `json:"Name"`
			Refer interface{} `json:"Refer"`
		} `json:"L"`
		R struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
		} `json:"R"`
	} `json:"Where"`
	Fields struct {
		Fields []struct {
			Offset   int         `json:"Offset"`
			WildCard interface{} `json:"WildCard"`
			Expr     struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
				Name struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Table struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Table"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
				} `json:"Name"`
				Refer interface{} `json:"Refer"`
			} `json:"Expr"`
			AsName struct {
				O string `json:"O"`
				L string `json:"L"`
			} `json:"AsName"`
			Auxiliary bool `json:"Auxiliary"`
		} `json:"Fields"`
	} `json:"Fields"`
	GroupBy struct {
		Items []struct {
			Expr struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
				Name struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Table struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Table"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
				} `json:"Name"`
				Refer interface{} `json:"Refer"`
			} `json:"Expr"`
			Desc      bool `json:"Desc"`
			NullOrder bool `json:"NullOrder"`
		} `json:"Items"`
	} `json:"GroupBy"`
	Having      interface{} `json:"Having"`
	WindowSpecs interface{} `json:"WindowSpecs"`
	OrderBy     struct {
		Items []struct {
			Expr struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
				N     int         `json:"N"`
				P     interface{} `json:"P"`
				Refer interface{} `json:"Refer"`
			} `json:"Expr"`
			Desc      bool `json:"Desc"`
			NullOrder bool `json:"NullOrder"`
		} `json:"Items"`
		ForUnion bool `json:"ForUnion"`
	} `json:"OrderBy"`
	Limit struct {
		Count struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
		} `json:"Count"`
		Offset struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
		} `json:"Offset"`
	} `json:"Limit"`
	LockInfo         interface{} `json:"LockInfo"`
	TableHints       interface{} `json:"TableHints"`
	IsInBraces       bool        `json:"IsInBraces"`
	WithBeforeBraces bool        `json:"WithBeforeBraces"`
	QueryBlockOffset int         `json:"QueryBlockOffset"`
	SelectIntoOpt    interface{} `json:"SelectIntoOpt"`
	AfterSetOperator interface{} `json:"AfterSetOperator"`
	Kind             int         `json:"Kind"`
	Lists            interface{} `json:"Lists"`
	With             interface{} `json:"With"`
}

/*
{
  "TableRefs": {
    "TableRefs": {
      "Left": {
        "Source": {
          "Schema": {
            "O": "",
            "L": ""
          },
          "Name": {
            "O": "t1",
            "L": "t1"
          },
          "DBInfo": null,
          "TableInfo": null,
          "IndexHints": [],
          "PartitionNames": [],
          "TableSample": null,
          "AsOf": null
        },
        "AsName": {
          "O": "",
          "L": ""
        }
      },
      "Right": null,
      "Tp": 0,
      "On": null,
      "Using": null,
      "NaturalJoin": false,
      "StraightJoin": false,
      "ExplicitParens": false
    }
  },
  "List": [
    {
      "Column": {
        "Schema": {
          "O": "",
          "L": ""
        },
        "Table": {
          "O": "",
          "L": ""
        },
        "Name": {
          "O": "a",
          "L": "a"
        }
      },
      "Expr": {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      }
    },
    {
      "Column": {
        "Schema": {
          "O": "",
          "L": ""
        },
        "Table": {
          "O": "",
          "L": ""
        },
        "Name": {
          "O": "b",
          "L": "b"
        }
      },
      "Expr": {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      }
    }
  ],
  "Where": {
    "Type": {
      "Tp": 0,
      "Flag": 0,
      "Flen": 0,
      "Decimal": 0,
      "Charset": "",
      "Collate": "",
      "Elems": null
    },
    "Op": 7,
    "L": {
      "Type": {
        "Tp": 0,
        "Flag": 0,
        "Flen": 0,
        "Decimal": 0,
        "Charset": "",
        "Collate": "",
        "Elems": null
      },
      "Name": {
        "Schema": {
          "O": "",
          "L": ""
        },
        "Table": {
          "O": "",
          "L": ""
        },
        "Name": {
          "O": "zz",
          "L": "zz"
        }
      },
      "Refer": null
    },
    "R": {
      "Type": {
        "Tp": 8,
        "Flag": 128,
        "Flen": 1,
        "Decimal": 0,
        "Charset": "binary",
        "Collate": "binary",
        "Elems": null
      }
    }
  },
  "Order": null,
  "Limit": null,
  "Priority": 0,
  "IgnoreErr": false,
  "MultipleTable": false,
  "TableHints": null,
  "With": null
}

*/

// UpDateStmt 更新操作结构体
type UpDateStmt struct {
	TableRefs struct {
		TableRefs struct {
			Left struct {
				Source struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
					DBInfo         interface{}   `json:"DBInfo"`
					TableInfo      interface{}   `json:"TableInfo"`
					IndexHints     []interface{} `json:"IndexHints"`
					PartitionNames []interface{} `json:"PartitionNames"`
					TableSample    interface{}   `json:"TableSample"`
					AsOf           interface{}   `json:"AsOf"`
				} `json:"Source"`
				AsName struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"AsName"`
			} `json:"Left"`
			Right          interface{} `json:"Right"`
			Tp             int         `json:"Tp"`
			On             interface{} `json:"On"`
			Using          interface{} `json:"Using"`
			NaturalJoin    bool        `json:"NaturalJoin"`
			StraightJoin   bool        `json:"StraightJoin"`
			ExplicitParens bool        `json:"ExplicitParens"`
		} `json:"TableRefs"`
	} `json:"TableRefs"`
	List []struct {
		Column struct {
			Schema struct {
				O string `json:"O"`
				L string `json:"L"`
			} `json:"Schema"`
			Table struct {
				O string `json:"O"`
				L string `json:"L"`
			} `json:"Table"`
			Name struct {
				O string `json:"O"`
				L string `json:"L"`
			} `json:"Name"`
		} `json:"Column"`
		Expr struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
		} `json:"Expr"`
	} `json:"List"`
	Where struct {
		Type struct {
			Tp      int         `json:"Tp"`
			Flag    int         `json:"Flag"`
			Flen    int         `json:"Flen"`
			Decimal int         `json:"Decimal"`
			Charset string      `json:"Charset"`
			Collate string      `json:"Collate"`
			Elems   interface{} `json:"Elems"`
		} `json:"Type"`
		Op int `json:"Op"`
		L  struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
			Name struct {
				Schema struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Schema"`
				Table struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Table"`
				Name struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"Name"`
			} `json:"Name"`
			Refer interface{} `json:"Refer"`
		} `json:"L"`
		R struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
		} `json:"R"`
	} `json:"Where"`
	Order         interface{} `json:"Order"`
	Limit         interface{} `json:"Limit"`
	Priority      int         `json:"Priority"`
	IgnoreErr     bool        `json:"IgnoreErr"`
	MultipleTable bool        `json:"MultipleTable"`
	TableHints    interface{} `json:"TableHints"`
	With          interface{} `json:"With"`
}

/*
{
  "TableRefs": {
    "TableRefs": {
      "Left": {
        "Source": {
          "Schema": {
            "O": "",
            "L": ""
          },
          "Name": {
            "O": "t1",
            "L": "t1"
          },
          "DBInfo": null,
          "TableInfo": null,
          "IndexHints": [],
          "PartitionNames": [],
          "TableSample": null,
          "AsOf": null
        },
        "AsName": {
          "O": "",
          "L": ""
        }
      },
      "Right": null,
      "Tp": 0,
      "On": null,
      "Using": null,
      "NaturalJoin": false,
      "StraightJoin": false,
      "ExplicitParens": false
    }
  },
  "Tables": null,
  "Where": {
    "Type": {
      "Tp": 0,
      "Flag": 0,
      "Flen": 0,
      "Decimal": 0,
      "Charset": "",
      "Collate": "",
      "Elems": null
    },
    "Op": 1,
    "L": {
      "Type": {
        "Tp": 0,
        "Flag": 0,
        "Flen": 0,
        "Decimal": 0,
        "Charset": "",
        "Collate": "",
        "Elems": null
      },
      "Op": 7,
      "L": {
        "Type": {
          "Tp": 0,
          "Flag": 0,
          "Flen": 0,
          "Decimal": 0,
          "Charset": "",
          "Collate": "",
          "Elems": null
        },
        "Name": {
          "Schema": {
            "O": "",
            "L": ""
          },
          "Table": {
            "O": "",
            "L": ""
          },
          "Name": {
            "O": "a",
            "L": "a"
          }
        },
        "Refer": null
      },
      "R": {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      }
    },
    "R": {
      "Type": {
        "Tp": 0,
        "Flag": 0,
        "Flen": 0,
        "Decimal": 0,
        "Charset": "",
        "Collate": "",
        "Elems": null
      },
      "Op": 7,
      "L": {
        "Type": {
          "Tp": 0,
          "Flag": 0,
          "Flen": 0,
          "Decimal": 0,
          "Charset": "",
          "Collate": "",
          "Elems": null
        },
        "Name": {
          "Schema": {
            "O": "",
            "L": ""
          },
          "Table": {
            "O": "",
            "L": ""
          },
          "Name": {
            "O": "b",
            "L": "b"
          }
        },
        "Refer": null
      },
      "R": {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      }
    }
  },
  "Order": null,
  "Limit": null,
  "Priority": 0,
  "IgnoreErr": false,
  "Quick": false,
  "IsMultiTable": false,
  "BeforeFrom": false,
  "TableHints": null,
  "With": null
}

*/

// DeleteStmt  删除操作结构体
type DeleteStmt struct {
	TableRefs struct {
		TableRefs struct {
			Left struct {
				Source struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
					DBInfo         interface{}   `json:"DBInfo"`
					TableInfo      interface{}   `json:"TableInfo"`
					IndexHints     []interface{} `json:"IndexHints"`
					PartitionNames []interface{} `json:"PartitionNames"`
					TableSample    interface{}   `json:"TableSample"`
					AsOf           interface{}   `json:"AsOf"`
				} `json:"Source"`
				AsName struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"AsName"`
			} `json:"Left"`
			Right          interface{} `json:"Right"`
			Tp             int         `json:"Tp"`
			On             interface{} `json:"On"`
			Using          interface{} `json:"Using"`
			NaturalJoin    bool        `json:"NaturalJoin"`
			StraightJoin   bool        `json:"StraightJoin"`
			ExplicitParens bool        `json:"ExplicitParens"`
		} `json:"TableRefs"`
	} `json:"TableRefs"`
	Tables interface{} `json:"Tables"`
	Where  struct {
		Type struct {
			Tp      int         `json:"Tp"`
			Flag    int         `json:"Flag"`
			Flen    int         `json:"Flen"`
			Decimal int         `json:"Decimal"`
			Charset string      `json:"Charset"`
			Collate string      `json:"Collate"`
			Elems   interface{} `json:"Elems"`
		} `json:"Type"`
		Op int `json:"Op"`
		L  struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
			Op int `json:"Op"`
			L  struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
				Name struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Table struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Table"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
				} `json:"Name"`
				Refer interface{} `json:"Refer"`
			} `json:"L"`
			R struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
			} `json:"R"`
		} `json:"L"`
		R struct {
			Type struct {
				Tp      int         `json:"Tp"`
				Flag    int         `json:"Flag"`
				Flen    int         `json:"Flen"`
				Decimal int         `json:"Decimal"`
				Charset string      `json:"Charset"`
				Collate string      `json:"Collate"`
				Elems   interface{} `json:"Elems"`
			} `json:"Type"`
			Op int `json:"Op"`
			L  struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
				Name struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Table struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Table"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
				} `json:"Name"`
				Refer interface{} `json:"Refer"`
			} `json:"L"`
			R struct {
				Type struct {
					Tp      int         `json:"Tp"`
					Flag    int         `json:"Flag"`
					Flen    int         `json:"Flen"`
					Decimal int         `json:"Decimal"`
					Charset string      `json:"Charset"`
					Collate string      `json:"Collate"`
					Elems   interface{} `json:"Elems"`
				} `json:"Type"`
			} `json:"R"`
		} `json:"R"`
	} `json:"Where"`
	Order        interface{} `json:"Order"`
	Limit        interface{} `json:"Limit"`
	Priority     int         `json:"Priority"`
	IgnoreErr    bool        `json:"IgnoreErr"`
	Quick        bool        `json:"Quick"`
	IsMultiTable bool        `json:"IsMultiTable"`
	BeforeFrom   bool        `json:"BeforeFrom"`
	TableHints   interface{} `json:"TableHints"`
	With         interface{} `json:"With"`
}

/*
{
  "IsReplace": false,
  "IgnoreErr": false,
  "Table": {
    "TableRefs": {
      "Left": {
        "Source": {
          "Schema": {
            "O": "",
            "L": ""
          },
          "Name": {
            "O": "test",
            "L": "test"
          },
          "DBInfo": null,
          "TableInfo": null,
          "IndexHints": null,
          "PartitionNames": null,
          "TableSample": null,
          "AsOf": null
        },
        "AsName": {
          "O": "",
          "L": ""
        }
      },
      "Right": null,
      "Tp": 0,
      "On": null,
      "Using": null,
      "NaturalJoin": false,
      "StraightJoin": false,
      "ExplicitParens": false
    }
  },
  "Columns": [
    {
      "Schema": {
        "O": "",
        "L": ""
      },
      "Table": {
        "O": "",
        "L": ""
      },
      "Name": {
        "O": "a",
        "L": "a"
      }
    },
    {
      "Schema": {
        "O": "",
        "L": ""
      },
      "Table": {
        "O": "",
        "L": ""
      },
      "Name": {
        "O": "b",
        "L": "b"
      }
    },
    {
      "Schema": {
        "O": "",
        "L": ""
      },
      "Table": {
        "O": "",
        "L": ""
      },
      "Name": {
        "O": "c",
        "L": "c"
      }
    },
    {
      "Schema": {
        "O": "",
        "L": ""
      },
      "Table": {
        "O": "",
        "L": ""
      },
      "Name": {
        "O": "d",
        "L": "d"
      }
    }
  ],
  "Lists": [
    [
      {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      },
      {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      },
      {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      },
      {
        "Type": {
          "Tp": 8,
          "Flag": 128,
          "Flen": 1,
          "Decimal": 0,
          "Charset": "binary",
          "Collate": "binary",
          "Elems": null
        }
      }
    ]
  ],
  "Setlist": null,
  "Priority": 0,
  "OnDuplicate": null,
  "Select": null,
  "TableHints": null,
  "PartitionNames": []
}
*/

// InsertStmt  插入操作结构体
type InsertStmt struct {
	IsReplace bool `json:"IsReplace"`
	IgnoreErr bool `json:"IgnoreErr"`
	Table     struct {
		TableRefs struct {
			Left struct {
				Source struct {
					Schema struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Schema"`
					Name struct {
						O string `json:"O"`
						L string `json:"L"`
					} `json:"Name"`
					DBInfo         interface{} `json:"DBInfo"`
					TableInfo      interface{} `json:"TableInfo"`
					IndexHints     interface{} `json:"IndexHints"`
					PartitionNames interface{} `json:"PartitionNames"`
					TableSample    interface{} `json:"TableSample"`
					AsOf           interface{} `json:"AsOf"`
				} `json:"Source"`
				AsName struct {
					O string `json:"O"`
					L string `json:"L"`
				} `json:"AsName"`
			} `json:"Left"`
			Right          interface{} `json:"Right"`
			Tp             int         `json:"Tp"`
			On             interface{} `json:"On"`
			Using          interface{} `json:"Using"`
			NaturalJoin    bool        `json:"NaturalJoin"`
			StraightJoin   bool        `json:"StraightJoin"`
			ExplicitParens bool        `json:"ExplicitParens"`
		} `json:"TableRefs"`
	} `json:"Table"`
	Columns []struct {
		Schema struct {
			O string `json:"O"`
			L string `json:"L"`
		} `json:"Schema"`
		Table struct {
			O string `json:"O"`
			L string `json:"L"`
		} `json:"Table"`
		Name struct {
			O string `json:"O"`
			L string `json:"L"`
		} `json:"Name"`
	} `json:"Columns"`
	Lists [][]struct {
		Type struct {
			Tp      int         `json:"Tp"`
			Flag    int         `json:"Flag"`
			Flen    int         `json:"Flen"`
			Decimal int         `json:"Decimal"`
			Charset string      `json:"Charset"`
			Collate string      `json:"Collate"`
			Elems   interface{} `json:"Elems"`
		} `json:"Type"`
	} `json:"Lists"`
	Setlist        interface{}   `json:"Setlist"`
	Priority       int           `json:"Priority"`
	OnDuplicate    interface{}   `json:"OnDuplicate"`
	Select         interface{}   `json:"Select"`
	TableHints     interface{}   `json:"TableHints"`
	PartitionNames []interface{} `json:"PartitionNames"`
}
