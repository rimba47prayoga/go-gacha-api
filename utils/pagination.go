package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
	"github.com/rimba47prayoga/go-gacha-api/models"
	"gorm.io/gorm"
)

type Pagination struct {
    Page         	int         `json:"page" default:"1"`
	PageSize		int			`json:"page_size" default:"10"`
    Sort         	string      `json:"sort"`
    TotalRows    	int64       `json:"total_rows"`    
    TotalPages   	int         `json:"total_pages"`   
    Rows         	interface{} `json:"rows"`  
}

func (p *Pagination) GetOffset() int {  
    return (p.GetPage() - 1) * p.GetLimit() 
}   

func (p *Pagination) GetLimit() int {   
    if p.PageSize == 0 {   
        p.PageSize = 10    
    }   
    return p.PageSize  
}

func (p *Pagination) GetPage() int {    
    if p.Page == 0 {    
        p.Page = 1  
    }   
    return p.Page   
}

func (p *Pagination) GetSort() string { 
    if p.Sort == "" {   
        p.Sort = "ID asc"  
    }   
    return p.Sort   
}

func InitPagination(ctx *gin.Context) *Pagination {
	pagination := Pagination{}
	defaults.SetDefaults(&pagination)

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	if page > 0 {
        pagination.Page = page
    }
    if pageSize > 0 {
        pagination.PageSize = pageSize
    }
	return &pagination
}


func Paginate(value interface{}, pagination *Pagination) func(db *gorm.DB) *gorm.DB {  
    var totalRows int64 
    models.DB.Model(value).Count(&totalRows)   
    pagination.TotalRows = totalRows 
    totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PageSize)))
    pagination.TotalPages = totalPages  
    return func(db *gorm.DB) *gorm.DB { 
        return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())   
    }   
}   
