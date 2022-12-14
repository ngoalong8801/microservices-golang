// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ProductID    int64 `gorm:"primaryKey"`
	ProductName  string
	Price        int32
	NumInventory int32
	Categories   []Category `gorm:"many2many:product_category;foreignKey:ProductID;joinForeignKey:ProductID;References:CategoryID;joinReferences:CategoryID"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
