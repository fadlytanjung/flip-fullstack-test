import { ChevronLeft, ChevronRight } from 'lucide-react';
import styles from './Pagination.module.css';

export interface PaginationMeta {
  total: number;
  count: number;
  per_page: number;
  current_page: number;
  total_pages: number;
}

interface PaginationProps {
  meta: PaginationMeta;
  onPageChange: (page: number) => void;
  loading?: boolean;
}

export function Pagination({ meta, onPageChange, loading = false }: PaginationProps) {
  const { current_page, total_pages, total, per_page } = meta;
  
  const startItem = (current_page - 1) * per_page + 1;
  const endItem = Math.min(current_page * per_page, total);

  const canGoPrevious = current_page > 1;
  const canGoNext = current_page < total_pages;

  // Generate page numbers with smart ellipsis
  const generatePageNumbers = () => {
    const pages: (number | string)[] = [];
    const totalPagesToShow = 5;
    const halfRange = Math.floor(totalPagesToShow / 2);

    if (total_pages <= 7) {
      // Show all pages if total is 7 or less
      for (let i = 1; i <= total_pages; i++) {
        pages.push(i);
      }
    } else {
      // Always show first page
      pages.push(1);

      if (current_page <= halfRange + 1) {
        // Near the start: 1, 2, 3, 4, 5, ..., 10
        for (let i = 2; i <= totalPagesToShow; i++) {
          pages.push(i);
        }
        pages.push('ellipsis-end');
        pages.push(total_pages);
      } else if (current_page >= total_pages - halfRange) {
        // Near the end: 1, ..., 6, 7, 8, 9, 10
        pages.push('ellipsis-start');
        for (let i = total_pages - totalPagesToShow + 1; i < total_pages; i++) {
          pages.push(i);
        }
        pages.push(total_pages);
      } else {
        // In the middle: 1, ..., 4, 5, 6, ..., 10
        pages.push('ellipsis-start');
        for (let i = current_page - 1; i <= current_page + 1; i++) {
          pages.push(i);
        }
        pages.push('ellipsis-end');
        pages.push(total_pages);
      }
    }

    return pages;
  };

  const pageNumbers = generatePageNumbers();

  return (
    <div className={styles.pagination}>
      <div className={styles.info}>
        Showing {startItem} to {endItem} of {total} entries
      </div>
      
      <div className={styles.controls}>
        <button
          className={styles.button}
          onClick={() => onPageChange(current_page - 1)}
          disabled={!canGoPrevious || loading}
        >
          <ChevronLeft size={16} />
          Previous
        </button>
        
        <div className={styles.pages}>
          {pageNumbers.map((page, index) => {
            if (typeof page === 'string') {
              // Render ellipsis
              return (
                <span key={`${page}-${index}`} className={styles.ellipsis}>
                  ...
                </span>
              );
            }

            return (
              <button
                key={page}
                className={`${styles.pageButton} ${current_page === page ? styles.active : ''}`}
                onClick={() => onPageChange(page)}
                disabled={loading}
              >
                {page}
              </button>
            );
          })}
        </div>
        
        <button
          className={styles.button}
          onClick={() => onPageChange(current_page + 1)}
          disabled={!canGoNext || loading}
        >
          Next
          <ChevronRight size={16} />
        </button>
      </div>
    </div>
  );
}
