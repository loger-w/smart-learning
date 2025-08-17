import { TbBook } from "react-icons/tb";

interface AuthHeaderProps {
  title: string;
  description: string;
}

export const AuthHeader = ({ title, description }: AuthHeaderProps) => {
  return (
    <div className="text-center mb-8">
      <div className="mx-auto w-16 h-16 text-white text-4xl bg-blue-600 rounded-2xl flex items-center justify-center mb-4">
        <TbBook />
      </div>
      <h1 className="text-3xl font-bold text-foreground mb-2">{title}</h1>
      <p className="text-muted-foreground">{description}</p>
    </div>
  );
};
