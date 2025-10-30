'use client';

import { useEffect } from 'react';
import { CheckCircle2, XCircle, AlertTriangle, Info, X } from 'lucide-react';
import styles from './Notification.module.css';

export type NotificationType = 'success' | 'error' | 'warning' | 'info';

export interface NotificationProps {
  id: string;
  type: NotificationType;
  title?: string;
  message: string;
  duration?: number;
  onClose: (id: string) => void;
}

export function Notification({ 
  id,
  type, 
  title, 
  message, 
  duration = 5000,
  onClose 
}: NotificationProps) {
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

  useEffect(() => {
    if (duration > 0) {
      const timer = setTimeout(() => {
        onClose(id);
      }, duration);

      return () => clearTimeout(timer);
    }
  }, [id, duration, onClose]);

  return (
    <div className={`${styles.notification} ${className}`} role="alert">
      <div className={styles.iconContainer}>
        <Icon className={styles.icon} size={20} />
      </div>
      <div className={styles.content}>
        <p className={styles.title}>{title || defaultTitle}</p>
        <p className={styles.message}>{message}</p>
      </div>
      <button
        className={styles.closeButton}
        onClick={() => onClose(id)}
        aria-label="Close notification"
      >
        <X size={18} />
      </button>
    </div>
  );
}

