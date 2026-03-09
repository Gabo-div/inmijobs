import { formatRelative } from "date-fns";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { ScrollArea } from "../ui/scroll-area";
import type { Job } from "@/models/jobs";

interface Props {
  job?: Job;
}

export default function JobDetailView({ job }: Props) {
  if (!job) return <div className="flex-1 p-10 text-gray-500">Selecciona un empleo para ver los detalles</div>;

  return (
    <ScrollArea className="flex-1 bg-white h-full border-l border-gray-200">
      <div className="p-6 pb-0">
        <div className="flex justify-between items-start">
          <h1 className="text-2xl font-bold text-gray-900">{job.title}</h1>
          <button className="p-2 hover:bg-gray-100 rounded-full text-gray-600">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
            </svg>
          </button>
        </div>

        <div className="mt-2 text-sm">
          <span className="text-gray-900 font-medium">Company</span>
          <span className="text-gray-500"> · {job.location} · {formatRelative(new Date(job.createdAt), new Date())}</span>
        </div>

        <div className="flex flex-wrap gap-2 mt-4">
          <Badge variant="secondary">
            {job.employmentType}
          </Badge>
          <Badge variant="secondary">
            ${job.salaryMin || "?"}
            {" "}
            -
            {" "}
            ${job.salaryMax || "?"}
          </Badge>
        </div>

        <div className="flex gap-2 mt-6">
          <Button type="submit" size="lg" className="w-full cursor-pointer font-bold bg-linear-to-r from-[#F97316] to-[#8B5CF6] text-white hover:from-[#EA580C] hover:to-[#7C3AED] shadow-lg transition-all duration-200">
            Solicitar
          </Button>
        </div>
      </div>

      <hr className="my-6 border-gray-200" />

      <div className="px-6 pb-10">
        <h2 className="text-xl font-bold text-gray-900 mb-4">Acerca del empleo</h2>
        <div className="text-sm text-gray-800 space-y-4 leading-relaxed whitespace-pre-line">
          {job.description}
          {job.description}
        </div>
      </div>
    </ScrollArea >
  );
};
