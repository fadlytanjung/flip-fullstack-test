'use client';

import { ReactNode } from 'react';
import { useSidebar } from '@/app/contexts';
import { Sidebar } from '../Sidebar';
import styles from './DashboardLayout.module.css';

interface DashboardLayoutProps {
  children: ReactNode;
}

export function DashboardLayout({ children }: DashboardLayoutProps) {
  const { collapsed } = useSidebar();
  
  // Use actual pixel values instead of CSS variables
  const sidebarWidth = collapsed ? 80 : 260;

  return (
    <div className={styles.dashboardLayout}>
      <Sidebar />
      <main 
        className={styles.mainContent}
        style={{
          width: `calc(100vw - ${sidebarWidth}px)`,
          marginLeft: `${sidebarWidth}px`
        }}
      >
        <div className={styles.contentWrapper}>
          {children}
        </div>
      </main>
    </div>
  );
}

