import { ReactNode } from 'react';
import styles from './ButtonGroup.module.css';

interface ButtonGroupProps {
  children: ReactNode;
  orientation?: 'horizontal' | 'vertical';
  gap?: 'small' | 'medium' | 'large';
  align?: 'start' | 'center' | 'end' | 'stretch';
}

export function ButtonGroup({
  children,
  orientation = 'horizontal',
  gap = 'medium',
  align = 'start'
}: ButtonGroupProps) {
  return (
    <div
      className={`${styles.buttonGroup} ${styles[orientation]} ${styles[`gap-${gap}`]} ${styles[`align-${align}`]}`}
    >
      {children}
    </div>
  );
}

