import { CheckCircle2, XCircle, AlertTriangle, Info, X } from 'lucide-react';
import styles from './Alert.module.css';

export type AlertType = 'success' | 'error' | 'warning' | 'info';

interface AlertProps {
  type: AlertType;
  title?: string;
  message: string;
  onClose?: () => void;
  dismissible?: boolean;
}

export function Alert({ 
  type, 
  title, 
  message, 
  onClose, 
  dismissible = false 
}: AlertProps) {
  const config = {
    success: {
      icon: CheckCircle2,
      title: title || 'Success',
      className: styles.success
    },
    error: {
      icon: XCircle,
      title: title || 'Error',
      className: styles.error
    },
    warning: {
      icon: AlertTriangle,
      title: title || 'Warning',
      className: styles.warning
    },
    info: {
      icon: Info,
      title: title || 'Info',
      className: styles.info
    }
  };

  const { icon: Icon, title: defaultTitle, className } = config[type];

  return (
    <div className={`${styles.alert} ${className}`} role="alert">
      <div className={styles.iconContainer}>
        <Icon className={styles.icon} size={20} />
      </div>
      <div className={styles.content}>
        <p className={styles.title}>{title || defaultTitle}</p>
        <p className={styles.message}>{message}</p>
      </div>
      {dismissible && onClose && (
        <button
          className={styles.closeButton}
          onClick={onClose}
          aria-label="Close alert"
        >
          <X size={18} />
        </button>
      )}
    </div>
  );
}

