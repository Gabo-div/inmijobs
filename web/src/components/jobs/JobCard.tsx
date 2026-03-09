import type { Job } from '@/models/jobs';

interface Props {
  job: Job;
  selected?: boolean;
  onSelect: (id: string) => void;
}

export function JobCard({ job, selected, onSelect }: Props) {
  return (
    <div
      onClick={() => onSelect(job.id)}
      className={`p-4 cursor-pointer transition-all ${selected ? 'w-full cursor-pointer bg-gray-100' : 'hover:bg-gray-100 bg-white'
        }`}
    >
      <div className="flex gap-4">
        <img src={job.company.logo} alt={""} className="size-14 rounded-full object-contain bg-white border border-gray-100" />
        <div className="flex-1">
          <h3 className="font-semibold hover:underline leading-tight">
            {job.title}
          </h3>
          <p className="text-sm text-gray-600">{job?.company.name}</p>
          <p className="text-xs mt-1 text-gray-500">{job.location}</p>
        </div>
      </div>
    </div>
  );
};
