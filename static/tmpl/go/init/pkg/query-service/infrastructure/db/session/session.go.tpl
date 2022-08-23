package session

import (
	"context"
	{{- if .IsMongo }}
	"{{.Namespace}}/pkg/query-service/infrastructure/db/dao/mongo_dao"
	{{- end }}
	{{- if .IsNeo4j }}
	"{{.Namespace}}/pkg/query-service/infrastructure/db/dao/neo4j_dao"
	{{- end }}
	"{{.Namespace}}/pkg/cmd-service/infrastructure/logs"
)

type Func func(ctx context.Context) error
type SessionType int

const (
	NoSession SessionType = iota
	ReadSession
	WriteSession
)


type Options struct {
    {{- if .IsMongo }}
    mongo *SessionType
    {{- end }}
    {{- if .IsNeo4j }}
	neo4j *SessionType
    {{- end }}
}

func NewOptions(opts ...*Options) *Options {
	return &Options{}
}

func MergeOptions(opts ...*Options) *Options {
	opt := &Options{}
	{{- if and .IsMongo .IsNeo4j }}
    w := WriteSession
	opt.mongo:= &w
	opt.neo4j:= &w
	{{- else if .IsMongo }}
    w := WriteSession
    opt.mongo = &w
	{{- else if .IsNeo4j }}
    w := WriteSession
    opt.neo4j = &w
	{{- end }}

	for _, o := range opts {
	    {{- if .IsMongo }}
		if o.mongo != nil {
			opt.mongo = o.mongo
		}
		{{- end }}
		{{- if .IsNeo4j }}
		if o.neo4j != nil {
			opt.neo4j = o.neo4j
		}
		{{- end }}
	}
	return opt
}

{{- if .IsMongo }}

func (o *Options) SetMongo(v SessionType) *Options {
    o.mongo = &v
    return o
}

func (o *Options) GetMongo() SessionType {
    if o.mongo == nil {
        return NoSession
    }
    return *o.mongo
}
{{- end }}

{{- if .IsNeo4j }}

func (o *Options) SetNeo4j(v SessionType) *Options {
	o.neo4j = &v
	return o
}

func (o *Options) GetNeo4j() SessionType {
	if o.neo4j == nil {
		return NoSession
	}
	return *o.neo4j
}
{{- end }}

func StartSession(ctx context.Context, fun Func, opts ...*Options) error {
	opt := MergeOptions(opts...)
    do := func() error {
        {{- if and .IsMongo .IsNeo4j }}
        if opt.GetMongo() != NoSession && opt.GetNeo4j() != NoSession {
            mongoSession := mongo_dao.NewSession(opt.GetMongo() == WriteSession)
            neo4jSession := neo4j_dao.NewSession(opt.GetNeo4j() == WriteSession)
            return mongoSession.UseTransaction(ctx, func(ctx context.Context) error {
                return neo4jSession.UseTransaction(ctx, func(ctx context.Context) error {
                    return fun(ctx)
                })
            })
        } else if opt.GetMongo() != NoSession {
            mongoSession := mongo_dao.NewSession(opt.GetMongo() == WriteSession)
            return mongoSession.UseTransaction(ctx, func(ctx context.Context) error {
                return fun(ctx)
            })
        } else if opt.GetNeo4j() != NoSession {
            neo4jSession := neo4j_dao.NewSession(opt.GetNeo4j() == WriteSession)
            return neo4jSession.UseTransaction(ctx, func(ctx context.Context) error {
                return fun(ctx)
            })
        }
        return fun(ctx)

        {{- else if .IsMongo }}
        if opt.GetMongo() != NoSession {
            mongoSession := mongo_dao.NewSession(opt.GetMongo() == WriteSession)
            return mongoSession.UseTransaction(ctx, func(ctx context.Context) error {
                return fun(ctx)
            })
        }
        return fun(ctx)

        {{- else if .IsNeo4j }}
        if opt.GetNeo4j() != NoSession {
            neo4jSession := neo4j_dao.NewSession(opt.GetNeo4j() == WriteSession)
            return neo4jSession.UseTransaction(ctx, func(ctx context.Context) error {
                return fun(ctx)
            })
        }
        return fun(ctx)
    {{- end}}
    }

    err := do()
	if err != nil {
        logs.Errorln("db.StartSession()", err)
    }
	return err
}
