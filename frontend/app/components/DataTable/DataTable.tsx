'use client';

import { useState } from 'react';
import { CheckCircle2, XCircle, AlertCircle, ArrowUpDown, Inbox } from 'lucide-react';
import styles from './DataTable.module.css';

export interface Transaction {
  id: string;
  timestamp: number;
  name: string;
  type: 'CREDIT' | 'DEBIT';
  amount: number;
  status: 'SUCCESS' | 'FAILED' | 'PENDING';
  description: string;
}

interface DataTableProps {
  transactions: Transaction[];
  title?: string;
  showFilters?: boolean;
}

export function DataTable({ transactions, title = 'Transactions', showFilters = true }: DataTableProps) {
  const [sortField, setSortField] = useState<keyof Transaction>('timestamp');
  const [sortDirection, setSortDirection] = useState<'asc' | 'desc'>('desc');
  const [filterStatus, setFilterStatus] = useState<string>('all');
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 10;

  const handleSort = (field: keyof Transaction) => {
    if (sortField === field) {
      setSortDirection(sortDirection === 'asc' ? 'desc' : 'asc');
    } else {
      setSortField(field);
      setSortDirection('asc');
    }
  };

  const filteredTransactions = transactions.filter(transaction => {
    if (filterStatus === 'all') return true;
    return transaction.status === filterStatus;
  });

  const sortedTransactions = [...filteredTransactions].sort((a, b) => {
    const aValue = a[sortField];
    const bValue = b[sortField];
    
    if (aValue < bValue) return sortDirection === 'asc' ? -1 : 1;
    if (aValue > bValue) return sortDirection === 'asc' ? 1 : -1;
    return 0;
  });

  const totalPages = Math.ceil(sortedTransactions.length / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const paginatedTransactions = sortedTransactions.slice(startIndex, startIndex + itemsPerPage);

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
    }).format(amount);
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getStatusClass = (status: string) => {
    switch (status) {
      case 'SUCCESS':
        return styles.statusSuccess;
      case 'FAILED':
        return styles.statusFailed;
      case 'PENDING':
        return styles.statusPending;
      default:
        return '';
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'SUCCESS':
        return <CheckCircle2 size={14} />;
      case 'FAILED':
        return <XCircle size={14} />;
      case 'PENDING':
        return <AlertCircle size={14} />;
      default:
        return null;
    }
  };

  return (
    <div className={styles.dataTable}>
      <div className={styles.header}>
        <div>
          <h2 className={styles.title}>{title}</h2>
          <p className={styles.subtitle}>{sortedTransactions.length} total transactions</p>
        </div>
        
        {showFilters && (
          <div className={styles.filters}>
            <label className={styles.filterLabel}>
              Filter by status:
              <select 
                value={filterStatus} 
                onChange={(e) => {
                  setFilterStatus(e.target.value);
                  setCurrentPage(1);
                }}
                className={styles.filterSelect}
              >
                <option value="all">All</option>
                <option value="SUCCESS">Success</option>
                <option value="FAILED">Failed</option>
                <option value="PENDING">Pending</option>
              </select>
            </label>
          </div>
        )}
      </div>

      <div className={styles.tableWrapper}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th onClick={() => handleSort('timestamp')} className={styles.sortable}>
                <span className={styles.headerContent}>
                  Date
                  {sortField === 'timestamp' && <ArrowUpDown size={14} className={styles.sortIcon} />}
                </span>
              </th>
              <th onClick={() => handleSort('name')} className={styles.sortable}>
                <span className={styles.headerContent}>
                  Name
                  {sortField === 'name' && <ArrowUpDown size={14} className={styles.sortIcon} />}
                </span>
              </th>
              <th onClick={() => handleSort('type')} className={styles.sortable}>
                <span className={styles.headerContent}>
                  Type
                  {sortField === 'type' && <ArrowUpDown size={14} className={styles.sortIcon} />}
                </span>
              </th>
              <th onClick={() => handleSort('amount')} className={styles.sortable}>
                <span className={styles.headerContent}>
                  Amount
                  {sortField === 'amount' && <ArrowUpDown size={14} className={styles.sortIcon} />}
                </span>
              </th>
              <th onClick={() => handleSort('status')} className={styles.sortable}>
                <span className={styles.headerContent}>
                  Status
                  {sortField === 'status' && <ArrowUpDown size={14} className={styles.sortIcon} />}
                </span>
              </th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {paginatedTransactions.length === 0 ? (
              <tr>
                <td colSpan={6} className={styles.emptyState}>
                  <Inbox className={styles.emptyIcon} size={48} />
                  <p>No transactions found</p>
                </td>
              </tr>
            ) : (
              paginatedTransactions.map((transaction) => (
                <tr key={transaction.id}>
                  <td className={styles.dateCell}>{formatDate(transaction.timestamp)}</td>
                  <td className={styles.nameCell}>{transaction.name}</td>
                  <td>
                    <span className={`${styles.badge} ${transaction.type === 'CREDIT' ? styles.badgeCredit : styles.badgeDebit}`}>
                      {transaction.type}
                    </span>
                  </td>
                  <td className={styles.amountCell}>
                    <span className={transaction.type === 'CREDIT' ? styles.amountCredit : styles.amountDebit}>
                      {transaction.type === 'CREDIT' ? '+' : '-'}{formatCurrency(transaction.amount)}
                    </span>
                  </td>
                  <td>
                    <span className={`${styles.status} ${getStatusClass(transaction.status)}`}>
                      <span className={styles.statusIcon}>{getStatusIcon(transaction.status)}</span>
                      {transaction.status}
                    </span>
                  </td>
                  <td className={styles.descriptionCell}>{transaction.description}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      {totalPages > 1 && (
        <div className={styles.pagination}>
          <button
            className={styles.paginationBtn}
            onClick={() => setCurrentPage(p => Math.max(1, p - 1))}
            disabled={currentPage === 1}
          >
            Previous
          </button>
          <div className={styles.paginationInfo}>
            Page {currentPage} of {totalPages}
          </div>
          <button
            className={styles.paginationBtn}
            onClick={() => setCurrentPage(p => Math.min(totalPages, p + 1))}
            disabled={currentPage === totalPages}
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
}

