'use client';

import { Notification, NotificationProps } from './Notification';
import styles from './NotificationContainer.module.css';

interface NotificationContainerProps {
  notifications: Omit<NotificationProps, 'onClose'>[];
  onClose: (id: string) => void;
  position?: 'top-right' | 'top-left' | 'bottom-right' | 'bottom-left' | 'top-center' | 'bottom-center';
}

export function NotificationContainer({ 
  notifications, 
  onClose,
  position = 'top-right' 
}: NotificationContainerProps) {
  return (
    <div className={`${styles.container} ${styles[position]}`}>
      {notifications.map((notification) => (
        <Notification
          key={notification.id}
          {...notification}
          onClose={onClose}
        />
      ))}
    </div>
  );
}

