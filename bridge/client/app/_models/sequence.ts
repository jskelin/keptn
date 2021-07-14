import {ResultTypes} from './result-types';
import {EVENT_ICONS} from "./event-icons";

export class Sequence {
  name: string;
  project: string;
  service: string;
  shkeptncontext: string;
  stages: [
    {
      image: string,
      latestEvaluation: {
        result: ResultTypes,
        score: number
      },
      latestEvent: {
        id: string,
        time: string,
        type: string
      },
      latestFailedEvent: {
        id: string,
        time: string,
        type: string
      },
      name: string
    }
  ];
  state: string;
  time: string;
  problemTitle?: string;

  icon: string;

  public isLoading() {
    return this.state === 'loading' || this.state === 'triggered';
  }

  public isFinished(): boolean {
    return this.state === 'finished';
  }

  public isSuccessful(): boolean {
    return this.isFinished();
  }

  public isProblem(): boolean {
    return this.name === 'problem'; // TODO: check this?!
  }

  public hasPendingApproval(): boolean {
    return this.getLatestEvent().type === 'approval'; // TODO: check this?!
  }

  public getIcon(): string {
    if(!this.icon) {
      this.icon = EVENT_ICONS[this.name] || EVENT_ICONS['default'];
    }
    return this.icon;
  }

  public getLabel(): string {
    return this.name;
  }

  public getStatusLabel(): string {
    return this.state;
  }

  public getLatestEvent() {
    return this.stages[this.stages.length-1].latestEvent; // TODO: return Trace
  }

  public getLastStage(): string {
    return this.stages[this.stages.length-1].name;
  }

  public getStage(stageName: string) {
    return this.stages.find(stage => stage.name === stageName);
  }

  public isFaulty(stageName: string) {
    return !!this.getStage(stageName).latestFailedEvent;
  }

  static fromJSON(data: any): Sequence {
    return Object.assign(new this(), data);
  }
}
