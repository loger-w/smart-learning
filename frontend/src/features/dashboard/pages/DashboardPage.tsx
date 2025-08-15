import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  BookOpen,
  CheckCircle,
  FileText,
  Plus,
  BarChart3,
  Zap,
} from "lucide-react";

const statsData = [
  {
    title: "單字清單",
    value: "0",
    icon: BookOpen,
    description: "已建立的學習清單",
    iconBg: "bg-blue-100 dark:bg-blue-900",
    iconColor: "text-blue-600 dark:text-blue-400",
  },
  {
    title: "已學習單字",
    value: "0",
    icon: CheckCircle,
    description: "累計學習完成",
    iconBg: "bg-green-100 dark:bg-green-900",
    iconColor: "text-green-600 dark:text-green-400",
  },
  {
    title: "學習天數",
    value: "0",
    icon: Zap,
    description: "持續學習記錄",
    iconBg: "bg-purple-100 dark:bg-purple-900",
    iconColor: "text-purple-600 dark:text-purple-400",
  },
];

export const DashboardPage = () => {
  return (
    <>
      <div className="mb-8">
        <h2 className="text-3xl font-bold text-foreground mb-2">儀表板</h2>
        <p className="text-muted-foreground">開始您的英語學習之旅</p>
      </div>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        {statsData.map((stat) => {
          const Icon = stat.icon;
          return (
            <Card key={stat.title} className="p-6">
              <div className="flex items-center">
                <div className={`w-12 h-12 ${stat.iconBg} rounded-lg flex items-center justify-center`}>
                  <Icon className={`w-6 h-6 ${stat.iconColor}`} />
                </div>
                <div className="ml-4">
                  <CardDescription className="text-sm font-medium">
                    {stat.title}
                  </CardDescription>
                  <CardTitle className="text-2xl font-bold">
                    {stat.value}
                  </CardTitle>
                </div>
              </div>
            </Card>
          );
        })}
      </div>

      {/* Quick Actions */}
      <Card className="p-6">
        <CardHeader className="pb-4 px-0">
          <CardTitle className="text-lg font-semibold">快速開始</CardTitle>
        </CardHeader>
        <CardContent className="px-0">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <Button
              className="h-20 flex flex-col items-center justify-center space-y-2"
              variant="outline"
            >
              <Plus className="w-6 h-6" />
              <span>建立單字清單</span>
            </Button>

            <Button
              className="h-20 flex flex-col items-center justify-center space-y-2"
              variant="outline"
            >
              <FileText className="w-6 h-6" />
              <span>開始複習</span>
            </Button>

            <Button
              className="h-20 flex flex-col items-center justify-center space-y-2"
              variant="outline"
            >
              <BarChart3 className="w-6 h-6" />
              <span>學習統計</span>
            </Button>
          </div>
        </CardContent>
      </Card>
    </>
  );
};
