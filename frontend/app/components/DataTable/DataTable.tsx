'use client';

import { ReactNode, useEffect, useRef } from 'react';
import { Table, TableColumn } from '../Table/Table';
import { Pagination, PaginationMeta } from '../Pagination/Pagination';
import { Search } from '../Search/Search';
import { Select } from '../Select/Select';
import { TableSkeleton } from '../Loader/Skeleton';
import styles from './DataTable.module.css';

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
  sortBy?: string | null;
  sortOrder?: 'asc' | 'desc' | null;
  actions?: ReactNode;
  showSearch?: boolean;
  showPagination?: boolean;
  emptyMessage?: string;
  filterType?: string;
  onFilterTypeChange?: (value: string) => void;
  filterStatus?: string;
  onFilterStatusChange?: (value: string) => void;
}

export function DataTable<T extends Record<string, any>>({
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
  emptyMessage,
  filterType = '',
  onFilterTypeChange,
  filterStatus = '',
  onFilterStatusChange
}: DataTableProps<T>) {
  const tableRef = useRef<HTMLDivElement>(null);

  // Prevent scroll to top when data changes
  useEffect(() => {
    if (tableRef.current) {
      const scrollTop = tableRef.current.scrollTop;
      // Preserve scroll position
      return () => {
        if (tableRef.current && scrollTop > 0) {
          tableRef.current.scrollTop = scrollTop;
        }
      };
    }
  }, [data]);

  return (
    <div className={styles.dataTable} ref={tableRef}>
      {/* Header */}
      {(title || subtitle || showSearch || actions || onFilterTypeChange || onFilterStatusChange) && (
        <div className={styles.header}>
          <div className={styles.headerText}>
            {title && <h2 className={styles.title}>{title}</h2>}
            {subtitle && <p className={styles.subtitle}>{subtitle}</p>}
          </div>

          <div className={styles.headerActions}>
            <div className={styles.filterControls}>
              {onFilterTypeChange && (
                <Select
                  options={[
                    { value: '', label: 'All Types' },
                    { value: 'CREDIT', label: 'Credit' },
                    { value: 'DEBIT', label: 'Debit' },
                  ]}
                  value={filterType}
                  onChange={onFilterTypeChange}
                  disabled={loading}
                />
              )}

              {onFilterStatusChange && (
                <Select
                  options={[
                    { value: '', label: 'All Statuses' },
                    { value: 'SUCCESS', label: 'Success' },
                    { value: 'FAILED', label: 'Failed' },
                    { value: 'PENDING', label: 'Pending' },
                  ]}
                  value={filterStatus}
                  onChange={onFilterStatusChange}
                  disabled={loading}
                />
              )}
            </div>

            {showSearch && onSearchChange && (
              <Search
                value={searchValue}
                onChange={onSearchChange}
                placeholder="Search..."
                debounceMs={300}
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

