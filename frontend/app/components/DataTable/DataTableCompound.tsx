'use client';

import { ReactNode } from 'react';
import { Table, TableColumn } from '../Table/Table';
import { Pagination, PaginationMeta } from '../Pagination/Pagination';
import { Search } from '../Search/Search';
import { TableSkeleton } from '../Loader/Skeleton';
import styles from './DataTableCompound.module.css';

export interface DataTableProps<T = any> {
  title?: string;
  subtitle?: string;
  columns: TableColumn<T>[];
  data: T[];
  meta?: PaginationMeta;
  loading?: boolean;
  searchValue?: string;
  onSearchChange?: (value: string) => void;
  onPageChange?: (page: number) => void;
  onSort?: (key: string) => void;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
  actions?: ReactNode;
  showSearch?: boolean;
  showPagination?: boolean;
  emptyMessage?: string;
}

export function DataTableCompound<T extends Record<string, any>>({
  title,
  subtitle,
  columns,
  data,
  meta,
  loading = false,
  searchValue = '',
  onSearchChange,
  onPageChange,
  onSort,
  sortBy,
  sortOrder,
  actions,
  showSearch = true,
  showPagination = true,
  emptyMessage
}: DataTableProps<T>) {
  
  return (
    <div className={styles.dataTable}>
      {/* Header */}
      {(title || subtitle || showSearch || actions) && (
        <div className={styles.header}>
          <div className={styles.headerText}>
            {title && <h2 className={styles.title}>{title}</h2>}
            {subtitle && <p className={styles.subtitle}>{subtitle}</p>}
          </div>
          
          <div className={styles.headerActions}>
            {showSearch && onSearchChange && (
              <Search
                value={searchValue}
                onChange={onSearchChange}
                placeholder="Search..."
              />
            )}
            {actions}
          </div>
        </div>
      )}

      {/* Table */}
      {loading ? (
        <TableSkeleton rows={5} columns={columns.length} />
      ) : (
        <Table
          columns={columns}
          data={data}
          loading={loading}
          onSort={onSort}
          sortBy={sortBy}
          sortOrder={sortOrder}
          emptyMessage={emptyMessage}
        />
      )}

      {/* Pagination */}
      {showPagination && meta && onPageChange && !loading && data.length > 0 && (
        <Pagination
          meta={meta}
          onPageChange={onPageChange}
          loading={loading}
        />
      )}
    </div>
  );
}

