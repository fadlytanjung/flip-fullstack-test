import { ReactNode } from 'react';
import { MoveUp, MoveDown, ArrowUpDown, InboxIcon } from 'lucide-react';
import styles from './Table.module.css';

export interface TableColumn<T = any> {
  key: string;
  header: string | ReactNode;
  render?: (row: T) => ReactNode;
  sortable?: boolean;
  width?: string;
}

interface TableProps<T = any> {
  columns: TableColumn<T>[];
  data: T[];
  loading?: boolean;
  emptyMessage?: string;
  onSort?: (key: string) => void;
  sortBy?: string | null;
  sortOrder?: 'asc' | 'desc' | null;
}

export function Table<T extends Record<string, any>>({
  columns,
  data,
  loading = false,
  emptyMessage = 'No data available',
  onSort,
  sortBy,
  sortOrder
}: TableProps<T>) {
  
  const handleSort = (key: string, sortable?: boolean) => {
    if (sortable && onSort) {
      onSort(key);
    }
  };

  const getSortIcon = (columnKey: string) => {
    if (sortBy !== columnKey || sortOrder === null) {
      return <ArrowUpDown size={16} className={styles.sortIconDefault} />;
    }

    if (sortOrder === 'asc') {
      return <MoveUp size={16} className={styles.sortIconAsc} />;
    }

    return <MoveDown size={16} className={styles.sortIconDesc} />;
  };

  return (
    <div className={styles.tableWrapper}>
      <table className={styles.table}>
        <thead>
          <tr>
            {columns.map((column) => (
              <th
                key={column.key}
                className={column.sortable ? styles.sortable : ''}
                onClick={() => handleSort(column.key, column.sortable)}
                style={{ width: column.width }}
              >
                <span className={styles.headerContent}>
                  {column.header}
                  {column.sortable && (
                    <div className={`${styles.sortIconWrapper} ${sortBy === column.key ? styles.active : ''}`}>
                      {getSortIcon(column.key)}
                    </div>
                  )}
                </span>
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {loading ? (
            <tr>
              <td colSpan={columns.length} className={styles.loading}>
                Loading...
              </td>
            </tr>
          ) : data.length === 0 ? (
            <tr>
                <td colSpan={columns.length} className={styles.emptyCell}>
                  <div className={styles.emptyState}>
                    <InboxIcon className={styles.emptyIcon} size={48} />
                    <p className={styles.emptyMessage}>{emptyMessage}</p>
                  </div>
              </td>
            </tr>
          ) : (
            data.map((row, rowIndex) => (
              <tr key={rowIndex}>
                {columns.map((column) => (
                  <td key={column.key}>
                    {column.render ? column.render(row) : row[column.key]}
                  </td>
                ))}
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}
