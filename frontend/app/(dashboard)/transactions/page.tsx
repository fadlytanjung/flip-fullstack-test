import { DashboardLayout } from '@/app/components';
import { UnderMaintenance } from '@/app/ui/UnderMaintenance';

export default function TransactionsPage() {
  return (
    <DashboardLayout>
      <UnderMaintenance 
        title="Transactions Page"
        message="The transactions page is currently under development. Please check back later."
      />
    </DashboardLayout>
  );
}

