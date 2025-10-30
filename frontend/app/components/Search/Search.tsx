import { Search as SearchIcon, X } from 'lucide-react';
import styles from './Search.module.css';

interface SearchProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
  onClear?: () => void;
}

export function Search({ value, onChange, placeholder = 'Search...', onClear }: SearchProps) {
  return (
    <div className={styles.search}>
      <SearchIcon className={styles.icon} size={18} />
      <input
        type="text"
        className={styles.input}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
      />
      {value && (
        <button
          className={styles.clearButton}
          onClick={() => {
            onChange('');
            onClear?.();
          }}
          aria-label="Clear search"
        >
          <X size={16} />
        </button>
      )}
    </div>
  );
}

