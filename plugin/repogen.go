package plugin

import (
	"fmt"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func (p *OrmPlugin) generateDefaultRepository(file *generator.FileDescriptor) {
	for _, message := range file.Messages() {
		if !getMessageOptions(message).GetOrmable() {
			continue
		}
		p.UsingGoImports("context")

		genAll := p.hasPrimaryKey(p.getOrmable(p.TypeName(message))) && p.hasIDField(message)
		typeName := p.TypeName(message)
		p.generateRepositoryInterface(typeName, genAll)
		p.generateRepositoryConstructor(typeName)
		p.generateDefaultRepositoryStruct(typeName)

		p.generateRepoCreateHandler(typeName)
		// FIXME: Temporary fix for Ormable objects that have no ID field but
		// have pk.
		if genAll {
			p.generateRepoReadHandler(typeName)
			p.generateRepoDeleteHandler(typeName)
			p.generateRepoDeleteSetHandler(typeName)
			p.generateRepoStrictUpdateHandler(typeName)
			p.generateRepoPatchHandler(typeName)
		}
		p.generateRepoListHandler(typeName)
	}
}

func (p *OrmPlugin) generateRepositoryInterface(typeName string, all bool) {
	ormable := p.getOrmable(typeName)
	p.P(`// `, typeName, `Repository is a default repository`)
	p.P(`type `, typeName, `Repository interface {`)

	p.P(`// Create`, typeName, ` executes a basic gorm create call`)
	p.P(`Create`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error)`)
	// TODO : list signature
	p.P(`// List`, typeName, ` executes a gorm list call`)
	listSign := fmt.Sprint(`List`, typeName, `(ctx context.Context`)
	if p.listHasFiltering(ormable) {
		listSign += fmt.Sprint(`, f `, `*`, p.Import(queryImport), `.Filtering`)
	}
	if p.listHasSorting(ormable) {
		listSign += fmt.Sprint(`, s `, `*`, p.Import(queryImport), `.Sorting`)
	}
	if p.listHasPagination(ormable) {
		listSign += fmt.Sprint(`, p `, `*`, p.Import(queryImport), `.Pagination`)
	}
	selectFields := p.listHasFieldSelection(ormable)
	if selectFields {
		listSign += fmt.Sprint(`, fs `, `*`, p.Import(queryImport), `.FieldSelection`)
	}
	listSign += fmt.Sprint(`) ([]*`, typeName, `, error)`)
	p.P(listSign)
	if !all {
		p.P(`}`)
		return
	}
	p.P(`// Read`, typeName, ` executes a basic gorm read call`)
	if p.readHasFieldSelection(ormable) {
		p.P(`Read`, typeName, `(ctx context.Context, in *`, typeName, `, fs *`, p.Import(queryImport), `.FieldSelection) (*`, typeName, `, error)`)
	} else {
		p.P(`Read`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error)`)
	}
	p.P(`// Delete`, typeName, ` executes a basic gorm delete call`)
	p.P(`Delete`, typeName, `(ctx context.Context, in *`, typeName, `) error`)

	p.P(`// Delete`, typeName, `Set executes a basic gorm delete set call`)
	p.P(`Delete`, typeName, `Set(ctx context.Context, in []*`, typeName, `) error`)

	p.P(`// StrictUpdate`, typeName, ` clears first level 1:many children and then executes a gorm update call`)
	p.P(`StrictUpdate`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error)`)

	p.P(`// Patch`, typeName, ` executes a basic gorm update call with patch behavior`)
	p.P(`Patch`, typeName, `(ctx context.Context, in *`,
		typeName, `, updateMask *`, p.Import(fmImport), `.FieldMask) (*`, typeName, `, error)`)

	p.P(`}`)
}
func (p *OrmPlugin) generateRepositoryConstructor(typeName string) {
	p.P(`func New`, typeName, `Repository(db *`, p.Import(tkgormImport), `.DB) (`, typeName, `Repository, error) {`)
	p.P(`if db == nil {`)
	p.P(`return nil, `, p.Import("errors"), `.New("db cannot be nil")`)
	p.P(`}`)
	p.P(`return &`, typeName, `Repository{db: db}, nil`)
	p.P(`}`)
}
func (p *OrmPlugin) generateDefaultRepositoryStruct(typeName string) {
	p.P(`// Default`, typeName, `Repository implements `, typeName, `Repository`)
	p.P(`type Default`, typeName, `Repository struct {`)
	p.P(`db *`, p.Import(gormImport), `.DB`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoCreateHandler(typeName string) {
	p.P(`func (r *Default`, typeName, `Repository) Create`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error) {`)
	p.P(`return DefaultCreate`, typeName, `(ctx, in, r.db)`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoReadHandler(typeName string) {
	ormable := p.getOrmable(typeName)
	// Different behavior if there is a
	if p.readHasFieldSelection(ormable) {
		p.P(`func (r *Default`, typeName, `Repository) Read`, typeName, `(ctx context.Context, in *`,
			typeName, `, fs *`, p.Import(queryImport), `.FieldSelection) (*`, typeName, `, error) {`)
		p.P(`return DefaultRead`, typeName, `(ctx, in, r.db, fs)`)
	} else {
		p.P(`func (r *Default`, typeName, `Repository) Read`, typeName, `(ctx context.Context, in *`,
			typeName, `) (*`, typeName, `, error) {`)
		p.P(`return DefaultRead`, typeName, `(ctx, in, r.db)`)
	}
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoDeleteHandler(typeName string) {
	p.P(`func (r *Default`, typeName, `Repository) Delete`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error) {`)
	p.P(`return DefaultDelete`, typeName, `(ctx, in, r.db)`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoDeleteSetHandler(typeName string) {
	p.P(`func (r *Default`, typeName, `Repository) Delete`, typeName, `Set(ctx context.Context, in []*`, typeName, `) (*`, typeName, `, error) {`)
	p.P(`return DefaultDelete`, typeName, `Set(ctx, in, r.db)`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoStrictUpdateHandler(typeName string) {
	p.P(`func (r *Default`, typeName, `Repository) StrictUpdate`, typeName, `(ctx context.Context, in *`, typeName, `) (*`, typeName, `, error) {`)
	p.P(`return DefaultStrictUpdate`, typeName, `(ctx, in, r.db)`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoPatchHandler(typeName string) {
	p.P(`func (r *Default`, typeName, `Repository) Patch`, typeName, `(ctx context.Context, in *`,
		typeName, `, updateMask *`, p.Import(fmImport), `.FieldMask) (*`, typeName, `, error) {`)
	p.P(`return DefaultPatch`, typeName, `(ctx, in, updateMask, r.db)`)
	p.P(`}`)
}

func (p *OrmPlugin) generateRepoListHandler(typeName string) {
	ormable := p.getOrmable(typeName)
	listSign := fmt.Sprint(`func (r *Default`, typeName, `Repository) List`, typeName, `(ctx context.Context`)
	var f, s, pg, fs string
	if p.listHasFiltering(ormable) {
		listSign += fmt.Sprint(`, f `, `*`, p.Import(queryImport), `.Filtering`)
		f = "f"
	} else {
		f = "nil"
	}
	if p.listHasSorting(ormable) {
		listSign += fmt.Sprint(`, s `, `*`, p.Import(queryImport), `.Sorting`)
		s = "s"
	} else {
		s = "nil"
	}
	if p.listHasPagination(ormable) {
		listSign += fmt.Sprint(`, p `, `*`, p.Import(queryImport), `.Pagination`)
		pg = "p"
	} else {
		pg = "nil"
	}
	selectFields := p.listHasFieldSelection(ormable)
	if selectFields {
		listSign += fmt.Sprint(`, fs `, `*`, p.Import(queryImport), `.FieldSelection`)
		fs = "fs"
	} else {
		fs = "nil"
	}
	listSign += fmt.Sprint(`) ([]*`, typeName, `, error) {`)
	p.P(listSign)
	p.P(`return DefaultList`, typeName, `(ctx, r.db, `, f, `, `, s, `, `, pg, `, `, fs, `)`)
	p.P(`}`)
}
