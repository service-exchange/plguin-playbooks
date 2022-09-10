// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {useSelector} from 'react-redux';

import {GlobalState} from '@service-exchange/types/store';
import {Team} from '@service-exchange/types/teams';

import {getChannelsNameMapInTeam} from 'service-exchange-redux/selectors/entities/channels';
import {getCurrentTeam, getTeam} from 'service-exchange-redux/selectors/entities/teams';

import {formatText, messageHtmlToComponent} from 'src/webapp_globals';

export const useDefaultMarkdownOptions = ({team, ...opts}: {team?: Maybe<Team | Team['id']>} & Record<string, any> = {}) => {
    const selectedTeam = useSelector((state: GlobalState) => {
        if (typeof team === 'string') {
            return getTeam(state, team);
        }
        return team ?? getCurrentTeam(state);
    });
    const channelNamesMap = useSelector((state: GlobalState) => selectedTeam && getChannelsNameMapInTeam(state, selectedTeam.id));

    return {
        singleline: false,
        atMentions: true,
        mentionHighlight: false,
        team: selectedTeam,
        channelNamesMap,
        ...opts,
    };
};

type Props = {
    value: string;
    teamId?: string;
    options?: Record<string, any>;
};

const FormattedMarkdown = ({
    value,
    options,
}: Props) => {
    const opts = useDefaultMarkdownOptions(options);
    return messageHtmlToComponent(formatText(value, opts), true, {});
};

export default FormattedMarkdown;
