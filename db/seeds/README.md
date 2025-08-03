# Database Seeds

This directory contains seed data for the inventory system database tables.

## Available Seeds

The following tables have seed data available:

1. **mata_uang** (currencies) - Common currencies including IDR, USD, EUR, JPY, SGD, CNY, MYR
2. **bc_documents** - BC (Bea Cukai) document types and codes
3. **warehouses** - Sample warehouses for different product categories
4. **buyers** - Sample buyer companies with complete address and NPWP
5. **suppliers** - Sample supplier companies with complete address and NPWP  
6. **products** - Sample products across all categories (Bahan Baku, Bahan Penolong, Mesin/Sparepart, Barang Jadi)
7. **saldo_awal** - Initial stock balances for products

## Usage

To run individual seeds:

```bash
# Seed specific table
go run main.go seed -table=mata_uang
go run main.go seed -table=bc_documents
go run main.go seed -table=warehouses
go run main.go seed -table=buyers
go run main.go seed -table=suppliers
go run main.go seed -table=products
go run main.go seed -table=saldo_awal

# Seed all tables at once
go run main.go seed -table=all

# Delete all seeded data
go run main.go seed -table=delete-all
```

## Seeding Order

When seeding all tables, the order is important due to foreign key constraints:

1. roles (existing)
2. mata_uang (currencies)
3. bc_documents
4. warehouses
5. buyers
6. suppliers
7. products
8. saldo_awal (depends on products)

## Data Overview

### Products (18 items)
- **Bahan Baku (5 items)**: Besi Beton, Semen Portland, Pasir Halus, Kerikil, Kawat Bendrat
- **Bahan Penolong (4 items)**: Cat Primer, Thinner, Lem Epoxy, Silikon Sealant
- **Mesin/Sparepart (4 items)**: Motor Listrik, Bearing, V-Belt, Gear Box
- **Barang Jadi (5 items)**: Panel Beton, Tiang Pancang, Paving Block, U-Ditch, Box Culvert

### Warehouses (7 items)
- 2 Bahan Baku warehouses
- 1 Bahan Penolong warehouse
- 2 Mesin/Sparepart warehouses
- 2 Barang Jadi warehouses

### Currencies (7 items)
- IDR, USD, EUR, JPY, SGD, CNY, MYR

### BC Documents (10 items)
- Various BC categories: BC 2.3, BC 2.5, BC 2.6.1, BC 2.7, BC 4.0, BC 4.1

### Buyers & Suppliers (8 each)
- Complete company information with addresses and NPWP numbers
- Spread across major Indonesian cities

## Notes

- All IDs use ULID format (26 characters)
- Foreign key constraints are respected in seeding order
- Realistic Indonesian business data is used for companies
- Product codes follow a logical pattern (BB-, BP-, MS-, BJ-)
- Initial stock quantities vary by product type and category