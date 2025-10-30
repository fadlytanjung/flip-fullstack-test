'use client';

import { createContext, useContext, useState, ReactNode, useCallback } from 'react';
import { NotificationContainer } from '../components/Notification/NotificationContainer';
import { NotificationProps, NotificationType } from '../components/Notification';

interface NotificationContextType {
  showNotification: (notification: Omit<NotificationProps, 'id' | 'onClose'>) => void;
  showSuccess: (message: string, title?: string) => void;
  showError: (message: string, title?: string) => void;
  showWarning: (message: string, title?: string) => void;
  showInfo: (message: string, title?: string) => void;
  clearNotifications: () => void;
}

const NotificationContext = createContext<NotificationContextType | undefined>(undefined);

export function NotificationProvider({ children }: { children: ReactNode }) {
  const [notifications, setNotifications] = useState<Omit<NotificationProps, 'onClose'>[]>([]);

  const removeNotification = useCallback((id: string) => {
    setNotifications(prev => prev.filter(n => n.id !== id));
  }, []);

  const showNotification = useCallback((notification: Omit<NotificationProps, 'id' | 'onClose'>) => {
    const id = `notification-${Date.now()}-${Math.random()}`;
    setNotifications(prev => [...prev, { ...notification, id }]);
  }, []);

  const showSuccess = useCallback((message: string, title?: string) => {
    showNotification({ type: 'success', message, title });
  }, [showNotification]);

  const showError = useCallback((message: string, title?: string) => {
    showNotification({ type: 'error', message, title });
  }, [showNotification]);

  const showWarning = useCallback((message: string, title?: string) => {
    showNotification({ type: 'warning', message, title });
  }, [showNotification]);

  const showInfo = useCallback((message: string, title?: string) => {
    showNotification({ type: 'info', message, title });
  }, [showNotification]);

  const clearNotifications = useCallback(() => {
    setNotifications([]);
  }, []);

  return (
    <NotificationContext.Provider
      value={{
        showNotification,
        showSuccess,
        showError,
        showWarning,
        showInfo,
        clearNotifications,
      }}
    >
      {children}
      <NotificationContainer notifications={notifications} onClose={removeNotification} />
    </NotificationContext.Provider>
  );
}

export function useNotification() {
  const context = useContext(NotificationContext);
  if (context === undefined) {
    throw new Error('useNotification must be used within a NotificationProvider');
  }
  return context;
}

