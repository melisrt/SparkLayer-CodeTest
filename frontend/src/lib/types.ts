export type TodoItem = {
  id: number;
  title: string;
  description: string;
  priority: "urgent" | "medium" | "low";
  completed: boolean;
};
