'use client';

import { usePathname } from 'next/navigation';
import Link from 'next/link';
import { 
  LayoutDashboard, 
  CreditCard, 
  Upload, 
  TrendingUp, 
  Settings, 
  ChevronLeft, 
  ChevronRight,
  Building2,
  User,
  Sparkles
} from 'lucide-react';
import { useSidebar } from '@/app/contexts';
import styles from './Sidebar.module.css';

interface NavItem {
  label: string;
  icon: React.ComponentType<{ size?: number; className?: string }>;
  path: string;
}

export function Sidebar() {
  const { collapsed, toggleSidebar } = useSidebar();
  const pathname = usePathname();
  
  const navItems: NavItem[] = [
    { label: 'Dashboard', icon: LayoutDashboard, path: '/' },
    { label: 'Transactions', icon: CreditCard, path: '/transactions' },
    { label: 'Upload', icon: Upload, path: '/upload' },
    { label: 'Reports', icon: TrendingUp, path: '/reports' },
    { label: 'Demo', icon: Sparkles, path: '/demo' },
    { label: 'Settings', icon: Settings, path: '/settings' },
  ];
  
  // Check if path is active
  const isActive = (path: string) => {
    if (path === '/') {
      return pathname === '/';
    }
    return pathname.startsWith(path);
  };

  return (
    <aside className={`${styles.sidebar} ${collapsed ? styles.collapsed : ''}`}>
      <div className={styles.header}>
        <div className={styles.logo}>
          <Building2 className={styles.logoIcon} size={24} />
          {!collapsed && <span className={styles.logoText}>BankViewer</span>}
        </div>
        <button 
          className={styles.toggleBtn}
          onClick={toggleSidebar}
          aria-label="Toggle sidebar"
        >
          {collapsed ? <ChevronRight size={18} /> : <ChevronLeft size={18} />}
        </button>
      </div>

      <nav className={styles.nav}>
        <ul className={styles.navList}>
          {navItems.map((item) => {
            const IconComponent = item.icon;
            const active = isActive(item.path);
            return (
              <li key={item.path}>
                <Link 
                  href={item.path}
                  className={`${styles.navItem} ${active ? styles.active : ''}`}
                  title={collapsed ? item.label : undefined}
                >
                  <IconComponent className={styles.navIcon} size={20} />
                  {!collapsed && <span className={styles.navLabel}>{item.label}</span>}
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>

      <div className={styles.footer}>
        <div className={styles.user}>
          <div className={styles.avatar}>
            <User size={20} />
          </div>
          {!collapsed && (
            <div className={styles.userInfo}>
              <div className={styles.userName}>Admin User</div>
              <div className={styles.userRole}>Administrator</div>
            </div>
          )}
        </div>
      </div>
    </aside>
  );
}

