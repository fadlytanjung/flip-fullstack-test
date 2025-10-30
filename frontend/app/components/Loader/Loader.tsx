import { Loader2 } from 'lucide-react';
import styles from './Loader.module.css';

interface LoaderProps {
  size?: 'small' | 'medium' | 'large';
  text?: string;
}

export function Loader({ size = 'medium', text }: LoaderProps) {
  const sizeMap = {
    small: 16,
    medium: 24,
    large: 48
  };

  return (
    <div className={`${styles.loader} ${styles[size]}`}>
      <Loader2 className={styles.spinner} size={sizeMap[size]} />
      {text && <span className={styles.text}>{text}</span>}
    </div>
  );
}

