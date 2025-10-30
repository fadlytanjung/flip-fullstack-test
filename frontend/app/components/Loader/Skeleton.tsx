import styles from './Skeleton.module.css';

interface SkeletonProps {
  width?: string | number;
  height?: string | number;
  variant?: 'text' | 'circular' | 'rectangular';
  className?: string;
}

export function Skeleton({ 
  width = '100%', 
  height = '1em', 
  variant = 'rectangular',
  className = '' 
}: SkeletonProps) {
  const style = {
    width: typeof width === 'number' ? `${width}px` : width,
    height: typeof height === 'number' ? `${height}px` : height,
  };

  return (
    <div 
      className={`${styles.skeleton} ${styles[variant]} ${className}`}
      style={style}
    />
  );
}

export function TableSkeleton({ rows = 5, columns = 6 }: { rows?: number; columns?: number }) {
  return (
    <div className={styles.tableSkeleton}>
      {/* Header */}
      <div className={styles.tableHeader}>
        {Array.from({ length: columns }).map((_, i) => (
          <Skeleton key={i} height={40} />
        ))}
      </div>
      {/* Rows */}
      {Array.from({ length: rows }).map((_, rowIndex) => (
        <div key={rowIndex} className={styles.tableRow}>
          {Array.from({ length: columns }).map((_, colIndex) => (
            <Skeleton key={colIndex} height={60} />
          ))}
        </div>
      ))}
    </div>
  );
}

export function CardSkeleton() {
  return (
    <div className={styles.cardSkeleton}>
      <Skeleton variant="circular" width={64} height={64} />
      <div className={styles.cardContent}>
        <Skeleton width="60%" height={16} />
        <Skeleton width="80%" height={32} />
      </div>
    </div>
  );
}

