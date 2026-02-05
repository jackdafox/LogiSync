# Project LogiSync: Backend Requirements & Technical Specification

## 1. Executive Summary
**LogiSync** is a high-concurrency inventory management engine designed for real-time stock synchronization across global warehouses. The goal is to provide a robust, low-latency API that prevents over-selling and provides a clear audit trail for every item movement.

---

## 2. Core Requirements

### **Functional Requirements**
* **Inventory Tracking:** Real-time visibility of SKU counts (Available, Reserved, Damaged).
* **Reservation System:** Temporary stock locks (15-minute TTL) to ensure items in a user's cart aren't sold to someone else.
* **Multi-Warehouse Support:** Ability to route orders to the nearest warehouse with available stock.
* **Audit Logging:** Every manual adjustment or automated sale must be recorded with a reason and timestamp.

### **Non-Functional Requirements**
* **Performance:** API responses must be under 100ms.
* **Reliability:** The system must handle race conditions where two users attempt to buy the last item simultaneously.
* **Scalability:** Stateless design to allow horizontal scaling via Docker/Kubernetes.

---

## 3. Technical Stack
* **Language:** Go 1.21+
* **Primary Database:** PostgreSQL (ACID compliant)
* **Caching/Locking:** Redis
* **Communication:** RESTful API (JSON)
* **Infrastructure:** Dockerized environment

---

## 4. Database Schema (SQL)

The following schema ensures data integrity and high-performance indexing.



```sql
-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 1. Warehouses Table
CREATE TABLE warehouses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    location_code VARCHAR(50) UNIQUE NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Products (SKUs) Table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sku VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    low_stock_threshold INTEGER DEFAULT 10,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. Inventory Levels
CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID REFERENCES warehouses(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    quantity_available INTEGER NOT NULL DEFAULT 0,
    quantity_reserved INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(warehouse_id, product_id)
);

-- 4. Stock Audit Logs
CREATE TABLE stock_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    inventory_id UUID REFERENCES inventory(id),
    change_amount INTEGER NOT NULL,
    reason VARCHAR(100), -- 'restock', 'sale', 'return', 'adjustment'
    reference_id VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_inventory_product ON inventory(product_id);
CREATE INDEX idx_stock_logs_inventory ON stock_logs(inventory_id);