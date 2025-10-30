import { DashboardLayout } from '@/app/components';
import { UnderMaintenance } from '@/app/ui/UnderMaintenance';

export default function ReportsPage() {
  return (
    <DashboardLayout>
      <UnderMaintenance 
        title="Reports Page"
        message="The reports page is currently under development. Please check back later."
      />
    </DashboardLayout>
  );
}

